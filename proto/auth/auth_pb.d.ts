import * as jspb from 'google-protobuf'



export class Some extends jspb.Message {
  getX(): number;
  setX(value: number): Some;

  getY(): number;
  setY(value: number): Some;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Some.AsObject;
  static toObject(includeInstance: boolean, msg: Some): Some.AsObject;
  static serializeBinaryToWriter(message: Some, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Some;
  static deserializeBinaryFromReader(message: Some, reader: jspb.BinaryReader): Some;
}

export namespace Some {
  export type AsObject = {
    x: number,
    y: number,
  }
}

export class Thing extends jspb.Message {
  getName(): string;
  setName(value: string): Thing;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Thing.AsObject;
  static toObject(includeInstance: boolean, msg: Thing): Thing.AsObject;
  static serializeBinaryToWriter(message: Thing, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Thing;
  static deserializeBinaryFromReader(message: Thing, reader: jspb.BinaryReader): Thing;
}

export namespace Thing {
  export type AsObject = {
    name: string,
  }
}

