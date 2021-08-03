/**
 * @fileoverview gRPC-Web generated client stub for auth
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as proto_auth_auth_pb from '../../proto/auth/auth_pb';


export class AuthClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: any; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodInfoGetJWKS = new grpcWeb.AbstractClientBase.MethodInfo(
    proto_auth_auth_pb.GetJWKSResponse,
    (request: proto_auth_auth_pb.GetJWKSRequest) => {
      return request.serializeBinary();
    },
    proto_auth_auth_pb.GetJWKSResponse.deserializeBinary
  );

  getJWKS(
    request: proto_auth_auth_pb.GetJWKSRequest,
    metadata: grpcWeb.Metadata | null): Promise<proto_auth_auth_pb.GetJWKSResponse>;

  getJWKS(
    request: proto_auth_auth_pb.GetJWKSRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: proto_auth_auth_pb.GetJWKSResponse) => void): grpcWeb.ClientReadableStream<proto_auth_auth_pb.GetJWKSResponse>;

  getJWKS(
    request: proto_auth_auth_pb.GetJWKSRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: proto_auth_auth_pb.GetJWKSResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/auth.Auth/GetJWKS',
        request,
        metadata || {},
        this.methodInfoGetJWKS,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/auth.Auth/GetJWKS',
    request,
    metadata || {},
    this.methodInfoGetJWKS);
  }

  methodInfoGetToken = new grpcWeb.AbstractClientBase.MethodInfo(
    proto_auth_auth_pb.GetTokenResponse,
    (request: proto_auth_auth_pb.GetTokenRequest) => {
      return request.serializeBinary();
    },
    proto_auth_auth_pb.GetTokenResponse.deserializeBinary
  );

  getToken(
    request: proto_auth_auth_pb.GetTokenRequest,
    metadata: grpcWeb.Metadata | null): Promise<proto_auth_auth_pb.GetTokenResponse>;

  getToken(
    request: proto_auth_auth_pb.GetTokenRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: proto_auth_auth_pb.GetTokenResponse) => void): grpcWeb.ClientReadableStream<proto_auth_auth_pb.GetTokenResponse>;

  getToken(
    request: proto_auth_auth_pb.GetTokenRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: proto_auth_auth_pb.GetTokenResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/auth.Auth/GetToken',
        request,
        metadata || {},
        this.methodInfoGetToken,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/auth.Auth/GetToken',
    request,
    metadata || {},
    this.methodInfoGetToken);
  }

  methodInfoValidateToken = new grpcWeb.AbstractClientBase.MethodInfo(
    proto_auth_auth_pb.ValidateTokenResponse,
    (request: proto_auth_auth_pb.ValidateTokenRequest) => {
      return request.serializeBinary();
    },
    proto_auth_auth_pb.ValidateTokenResponse.deserializeBinary
  );

  validateToken(
    request: proto_auth_auth_pb.ValidateTokenRequest,
    metadata: grpcWeb.Metadata | null): Promise<proto_auth_auth_pb.ValidateTokenResponse>;

  validateToken(
    request: proto_auth_auth_pb.ValidateTokenRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: proto_auth_auth_pb.ValidateTokenResponse) => void): grpcWeb.ClientReadableStream<proto_auth_auth_pb.ValidateTokenResponse>;

  validateToken(
    request: proto_auth_auth_pb.ValidateTokenRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: proto_auth_auth_pb.ValidateTokenResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/auth.Auth/ValidateToken',
        request,
        metadata || {},
        this.methodInfoValidateToken,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/auth.Auth/ValidateToken',
    request,
    metadata || {},
    this.methodInfoValidateToken);
  }

}

