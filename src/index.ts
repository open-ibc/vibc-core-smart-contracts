import path from 'path'
import fs from 'fs'
import os from 'os'
import tar from 'tar'
import { contractsTemplate } from './contracts.template'

export async function extractVibcCoreSmartContracts(contractsDir: string): Promise<void> {
  // Create a unique temporary directory
  const tempDir = fs.mkdtempSync(os.tmpdir() + path.sep)

  // Write the decoded template to a temporary tar file
  const tempTarPath = path.join(tempDir, 'temp.tar')

  try {
    // Decode the base64-encoded contracts template
    const decodedTemplate = Buffer.from(contractsTemplate.trim(), 'base64')

    fs.writeFileSync(tempTarPath, decodedTemplate, 'binary')

    // Extract the tar file contents into the contracts directory
    await tar.extract({
      strip: 2,
      file: tempTarPath,
      cwd: contractsDir
    })
  } finally {
    // Delete the temporary directory
    fs.rmSync(tempDir, { force: true, recursive: true })
  }
}
