import {
  Command,
  Option,
  InvalidArgumentError,
} from "@commander-js/extra-typings";
import fs from "fs";
import path from "path";
import yaml from "yaml";
import * as winston from "winston";
import { Writable } from "stream";

export { Command, Option, InvalidArgumentError };
export type CmdArgsType<T extends (...args: any[]) => any> = Parameters<
  ReturnType<T>["handler"]
>[0];

/**
 * Extract type of a Command's action function
 */
export type Action<T extends (...params: any[]) => Command<any, any>> =
  Parameters<ReturnType<T>["action"]>[0];

/**
 * Extract params type of a Command's action function
 */
export type ActionParams<T extends (...params: any[]) => Command<any, any>> =
  Parameters<Action<T>>;

/**
 * Arg/Opt: an array of strings, split by comma. Whitespaces are not ignored.
 */
export function CommaStrs(value: string) {
  return value.split(",");
}

/**
 * Arg/Opt: an array of strings, split by comma or whitespace. Whitespaces are ignored.
 */
export function CommaOrWhitespaceStrs(value: string) {
  return value.split(/[, ]+/);
}

/**
 * Arg/Opt: an integer. NaN causes error.
 */
export function Int(value: string) {
  const parsed = parseInt(value);
  if (isNaN(parsed)) {
    throw new InvalidArgumentError(`invalid integer: ${value}`);
  }
  return parsed;
}

/**
 * Arg/Opt: a port integer [0, 65535]
 */
export function Port(value: string) {
  const parsed = Int(value);
  if (parsed < 0 || parsed > 65535) {
    throw new InvalidArgumentError(
      `invalid port: ${value}, must be [0, 65535]`
    );
  }
  return parsed;
}

/**
 * Arg/Opt: a parsed yaml object
 */
export function YamlFile(value: string) {
  const file = FilePath(value);
  return yaml.parse(fs.readFileSync(file, "utf8"));
}

/**
 * Arg/Opt: an existing file path
 */
export function FilePath(value: string) {
  if (!fs.existsSync(value)) {
    throw new InvalidArgumentError(`file not found: ${value}`);
  }
  if (!fs.statSync(value).isFile()) {
    throw new InvalidArgumentError(`not a file: ${value}`);
  }
  return value;
}

/**
 * Arg/Opt: an existing directory path
 */
export function DirPath(value: string) {
  if (!fs.existsSync(value)) {
    throw new InvalidArgumentError(`dir not found: ${value}`);
  }
  if (!fs.statSync(value).isDirectory()) {
    throw new InvalidArgumentError(`not a dir: ${value}`);
  }
  return value;
}

const OUTPUT_LOGGER_NAME = "output";
const MAIN_LOGGER_NAME = "main";

export const DEFAULT_OUTPUT = "-.yaml";
export const DEFAULT_LOGGER = "info:-";
export const MEM_TRANSPORT_NAME = "__mem__";

export type Logger = winston.Logger;

/**
 * returns the output logger
 * It returns a default logger (defined by DEFAULT_OUTPUT) if not set by user's cli option
 * Must use 'info' level for output logger
 */
export function getOutputLogger(): Logger {
  if (!winston.loggers.has(OUTPUT_LOGGER_NAME)) {
    OutputTarget(DEFAULT_OUTPUT);
  }
  return winston.loggers.get(OUTPUT_LOGGER_NAME);
}

/**
 * returns the main logger
 */
export function getMainLogger(): Logger {
  if (!winston.loggers.has(MAIN_LOGGER_NAME)) {
    SetMainLogger(DEFAULT_LOGGER);
  }
  return winston.loggers.get(MAIN_LOGGER_NAME);
}

/**
 * Arg/Opt: an output target formatted as <base>.<ext>, where base is a file path or "-" for stdout, and ext is a file
 * extension that determines the output format.
 */
export function OutputTarget(value: string) {
  // remove logger if already set
  if (winston.loggers.has(OUTPUT_LOGGER_NAME)) {
    winston.loggers.close(OUTPUT_LOGGER_NAME);
  }

  const valildExtensions = ["json", "yaml", "yml"];
  const parsed = path.parse(value);
  const base = path.join(parsed.dir, parsed.name);
  const ext = parsed.ext.length > 0 ? parsed.ext.slice(1) : "";
  if (!ext || valildExtensions.indexOf(ext) < 0) {
    throw new InvalidArgumentError(
      `invalid output extension: ${ext}, must be one of [${valildExtensions.join(
        ","
      )}]`
    );
  }
  let outFmt: winston.Logform.Format;
  // log the message only; no level or timestamp
  switch (ext) {
    case "json":
      outFmt = winston.format.printf((info: any) =>
        JSON.stringify(info.message, null, 2)
      );
      break;
    case "yaml":
    case "yml":
      outFmt = winston.format.printf((info) => yaml.stringify(info.message));
      break;
    default:
      throw new InvalidArgumentError(
        `invalid output extension: ${ext}, must be one of [${valildExtensions.join(
          ","
        )}]`
      );
  }
  let transport: winston.LoggerOptions["transports"];

  switch (base) {
    case "-":
      transport = new winston.transports.Console({});
      break;
    case MEM_TRANSPORT_NAME:
      transport = new winston.transports.Stream({ stream: new MemStream() });
      break;
    default:
      transport = new winston.transports.File({ filename: value, options: {} });
      break;
  }
  winston.loggers.add(OUTPUT_LOGGER_NAME, {
    transports: transport,
    format: outFmt,
    level: "info",
  });
  return value;
}

export function getOutputInMem() {
  const logger = getOutputLogger();
  const transport: { _stream: MemStream } = logger.transports[0] as any;
  if (
    !transport._stream ||
    typeof transport._stream.getContent() !== "string"
  ) {
    throw new Error("output logger is not in memory");
  }
  return transport._stream.getContent();
}

class MemStream extends Writable {
  _content: string = "";

  _write(
    chunk: any,
    encoding: BufferEncoding,
    callback: (error?: Error | null | undefined) => void
  ): void {
    this._content += chunk;
    if (callback) {
      callback();
    }
  }

  getContent() {
    return this._content;
  }
}

// from 0 to 6
const LogLevel = [
  "error",
  "warn",
  // default to info
  "info",
  "http",
  "verbose",
  "debug",
  "silly",
];

/**
 * parse a logger from a cli arg string, which may contain one or more logger transports, separated by comma
 * Logger can be later accessed by calling `getMainLogger` or `winston.loggers.get(MAIN_LOGGER_NAME)`
 * @param value a string of logger transports, separated by comma, eg. "info:-,error:err.log"
 * by default, file transport will overwrite existing file; use '+' to append to existing file, eg. "info:-,debug:+debug.log"
 * @returns same string as value
 */
export function SetMainLogger(value: string): string {
  const transports = parseLoggerTransports(value);
  // normally we wouldn't set logger multiple times;
  // for testing, we clean up existing logger before setting a new one
  if (winston.loggers.has(MAIN_LOGGER_NAME)) {
    winston.loggers.close(MAIN_LOGGER_NAME);
  }
  winston.loggers.add(MAIN_LOGGER_NAME, {
    transports,
    format: winston.format.combine(
      // winston.format.timestamp({ format: 'YYYY-MM-DD HH:mm:ss.SSS', options: { utc: true } }), // Use UTC time
      winston.format.timestamp(), // UTC time by default
      winston.format.printf((info) => {
        return `${info.timestamp} ${info.level}: ${info.message}`;
      })
    ),
  });
  return value;
}
/**
 * parse a logger from a arg string, which may contain one or more logger transports, separated by comma
 * Logger can be later accessed by calling `getMainLogger` or `winston.loggers.get(MAIN_LOGGER_NAME)`
 * @param value a string of logger transports, separated by comma, eg. "info:-,error:err.log"
 * @returns a winston logger
 */
export function createLogger(value: string) {
  return winston.createLogger({
    transports: parseLoggerTransports(value),
    format: winston.format.combine(
      // winston.format.timestamp({ format: 'YYYY-MM-DD HH:mm:ss.SSS', options: { utc: true } }), // Use UTC time
      winston.format.timestamp(), // UTC time by default
      winston.format.printf((info) => {
        return `${info.timestamp} ${info.level}: ${info.message}`;
      })
    ),
  });
}

/**
 * Create or use a logger from a string or use an existing Logger instance
 * @param value a string of logger transports, separated by comma, eg. "info:-,error:err.log";
 * Or a logger instance created somewhere else.
 * If undefined, use the `DEFAULT_LOGGER` string
 * @returns a new or existing logger
 */
export function createOrUseLogger(value?: string | Logger) {
  value = value || DEFAULT_LOGGER;
  if (typeof value === "string") {
    return createLogger(value);
  } else if (value instanceof winston.Logger) {
    return value;
  }
  throw new Error(
    `invalid logger, should be a string or Logger instance, but got: ${value}`
  );
}

function parseLoggerTransports(value: string) {
  const transportStrs = value.split(/,\s*/);
  const transports = transportStrs.map((str) => {
    // trim whitespaces
    const [level, filename] = str.split(/:\s*/);
    if (!LogLevel.includes(level))
      throw new Error(`invalid log level: ${level}`);
    if (!filename || filename === "-") {
      return new winston.transports.Console({ level });
    } else {
      // append to existing file
      if (filename.startsWith("+")) {
        return new winston.transports.File({
          filename: filename.slice(1),
          level,
          options: { flags: "a" },
        });
      }
      // overwrite existing file
      return new winston.transports.File({
        filename,
        level,
        options: { flags: "w" },
      });
    }
  });
  return transports;
}
