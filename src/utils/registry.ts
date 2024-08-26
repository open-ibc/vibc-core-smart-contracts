import assert from "assert";

/**
 * A template for a registry that maps strings to items of type T.
 * Items are ordered by insertion order.
 * Invariants:
 *  - keys are unique
 */
export class Registry<T, S = T> {
  public readonly registry: Map<string, T>;

  protected nameFunc: (t: T) => string;
  protected itemToObj: (t: T) => S;
  protected nameInParent?: string;

  /**
   * Create a new registry from a list of items.
   * @param src The source array of registry items
   * @param nameFunc A function that takes an item and returns its name, which must be unique.
   * Names are used for reigstry lookups.
   * Duplicated names result in the last item with the name being used.
   */
  constructor(
    src: T[],
    options?: {
      nameInParent?: string;
      nameFunc?: (t: T) => string;
      toObj?: (t: T) => S;
    }
  ) {
    const { nameFunc = (t: any) => t["name"], toObj = (t: any) => t } = options
      ? options
      : {};
    this.registry = new Map<string, T>();
    src.forEach((item) => {
      this.set(nameFunc(item), item);
    });
    this.itemToObj = toObj;
    this.nameFunc = nameFunc;
    this.nameInParent = options?.nameInParent;
  }

  /**
   * Get the element type of the registry.
   * Should only use this for type checking.
   * Returned run-time value is `undefined`.
   */
  get Element(): T {
    return undefined as any;
  }

  get name(): string {
    assert(this.nameInParent, "nameInParent not set");
    return this.nameInParent;
  }

  public get size(): number {
    return this.registry.size;
  }

  public keys(): string[] {
    return Array.from(this.registry.keys());
  }

  public values(): T[] {
    return Array.from(this.registry.values());
  }

  /**
   * Get an item by name. If the item doesn't exist, undefined is returned.
   */
  public get(name: string): T | undefined {
    return this.registry.get(name);
  }

  /**
   * Get an non-nullable item by name. If the item doesn't exist, an error is thrown.
   */
  public mustGet(name: string): T {
    const item = this.get(name);
    assert(item, `item ${name} not found in registry keys: ${this.keys()}`);
    return item;
  }

  /**
   * returns an iterator over the registry item tuples [key, value:T].
   */
  public entries(): IterableIterator<[string, T]> {
    return this.registry.entries();
  }

  /**
   * Copy items from another registry to this registry.
   * @param entries
   * @param allowDup If true, dup keys are allowed, ie. the last item with the same key is used.
   */
  public copyFrom(
    entries: IterableIterator<[string, T]>,
    allowDup = false
  ): void {
    for (const [name, item] of entries) {
      this.set(name, item, allowDup);
    }
  }

  /**
   * Set an item in the registry. If the item already exists, an error is thrown, unless force is true.
   */
  public set(name: string, item: T, force = false): void {
    if (!force && this.registry.has(name)) {
      throw new Error(`item ${name} already exists in registry`);
    }
    this.registry.set(name, item);
  }

  public has(name: string): boolean {
    return this.registry.has(name);
  }

  /**
   * Create a new registry instance from a subset of items in the current registry.
   * Error is thrown if any of the items in names is not found in the current registry, unless ignoreErr is true.
   * NOTE: Items in the new registry are shallow copies of the original items.
   * @param names items to include in the new registry.
   * If undefined, all items are included
   * If empty array, an empty registry is returned.
   * @param ignoreErr siliently ignore items not found in the current registry
   * @returns
   */
  public subset(names?: string[], ignoreErr = false): this {
    const classConstrutor = this.constructor as new (...args: any[]) => this;
    const subset = new classConstrutor([], {
      toObj: this.itemToObj,
      nameFunc: this.nameFunc,
      nameInParent: this.nameInParent,
    });
    const keys = names ? names : this.keys();
    for (const name of keys) {
      if (this.has(name)) {
        subset.set(name, this.mustGet(name), true);
      } else if (!ignoreErr) {
        throw new Error(
          `item ${name} not found in registry keys: ${this.keys()}`
        );
      }
    }
    return subset;
  }

  /**
   * @returns A list of items in the registry, ordered by insertion order.
   */
  public toList(): T[] {
    return Array.from(this.registry.values());
  }

  /**
   * Serialize the registry to a list of items. Default to toList().
   * Subclasses should override this method if they want to serialize the registry differently.
   * @returns A list of items in the registry, ordered by insertion order.
   */
  public serialize(): S[] {
    return this.toList().map((item) => this.itemToObj(item));
  }

  /**
   *
   * @returns An object with keys as the item names and values as the items.
   */
  toSerializedObj(): Record<string, S> {
    const obj: Record<string, S> = {};
    this.registry.forEach((item, key) => {
      obj[key] = this.itemToObj(item);
    });
    return obj;
  }
}
