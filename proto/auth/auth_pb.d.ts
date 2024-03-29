import * as jspb from 'google-protobuf'



export class GetJWKSRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetJWKSRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetJWKSRequest): GetJWKSRequest.AsObject;
  static serializeBinaryToWriter(message: GetJWKSRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetJWKSRequest;
  static deserializeBinaryFromReader(message: GetJWKSRequest, reader: jspb.BinaryReader): GetJWKSRequest;
}

export namespace GetJWKSRequest {
  export type AsObject = {
  }
}

export class GetJWKSResponse extends jspb.Message {
  getJwks(): string;
  setJwks(value: string): GetJWKSResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetJWKSResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetJWKSResponse): GetJWKSResponse.AsObject;
  static serializeBinaryToWriter(message: GetJWKSResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetJWKSResponse;
  static deserializeBinaryFromReader(message: GetJWKSResponse, reader: jspb.BinaryReader): GetJWKSResponse;
}

export namespace GetJWKSResponse {
  export type AsObject = {
    jwks: string,
  }
}

export class GetTokenRequest extends jspb.Message {
  getId(): string;
  setId(value: string): GetTokenRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetTokenRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetTokenRequest): GetTokenRequest.AsObject;
  static serializeBinaryToWriter(message: GetTokenRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetTokenRequest;
  static deserializeBinaryFromReader(message: GetTokenRequest, reader: jspb.BinaryReader): GetTokenRequest;
}

export namespace GetTokenRequest {
  export type AsObject = {
    id: string,
  }
}

export class GetTokenResponse extends jspb.Message {
  getToken(): string;
  setToken(value: string): GetTokenResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetTokenResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetTokenResponse): GetTokenResponse.AsObject;
  static serializeBinaryToWriter(message: GetTokenResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetTokenResponse;
  static deserializeBinaryFromReader(message: GetTokenResponse, reader: jspb.BinaryReader): GetTokenResponse;
}

export namespace GetTokenResponse {
  export type AsObject = {
    token: string,
  }
}

export class ValidateTokenRequest extends jspb.Message {
  getToken(): string;
  setToken(value: string): ValidateTokenRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ValidateTokenRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ValidateTokenRequest): ValidateTokenRequest.AsObject;
  static serializeBinaryToWriter(message: ValidateTokenRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ValidateTokenRequest;
  static deserializeBinaryFromReader(message: ValidateTokenRequest, reader: jspb.BinaryReader): ValidateTokenRequest;
}

export namespace ValidateTokenRequest {
  export type AsObject = {
    token: string,
  }
}

export class ValidateTokenResponse extends jspb.Message {
  getValid(): boolean;
  setValid(value: boolean): ValidateTokenResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ValidateTokenResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ValidateTokenResponse): ValidateTokenResponse.AsObject;
  static serializeBinaryToWriter(message: ValidateTokenResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ValidateTokenResponse;
  static deserializeBinaryFromReader(message: ValidateTokenResponse, reader: jspb.BinaryReader): ValidateTokenResponse;
}

export namespace ValidateTokenResponse {
  export type AsObject = {
    valid: boolean,
  }
}

