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

  methodInfoGetSomething = new grpcWeb.AbstractClientBase.MethodInfo(
    proto_auth_auth_pb.Thing,
    (request: proto_auth_auth_pb.Some) => {
      return request.serializeBinary();
    },
    proto_auth_auth_pb.Thing.deserializeBinary
  );

  getSomething(
    request: proto_auth_auth_pb.Some,
    metadata: grpcWeb.Metadata | null): Promise<proto_auth_auth_pb.Thing>;

  getSomething(
    request: proto_auth_auth_pb.Some,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: proto_auth_auth_pb.Thing) => void): grpcWeb.ClientReadableStream<proto_auth_auth_pb.Thing>;

  getSomething(
    request: proto_auth_auth_pb.Some,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: proto_auth_auth_pb.Thing) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/auth.Auth/GetSomething',
        request,
        metadata || {},
        this.methodInfoGetSomething,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/auth.Auth/GetSomething',
    request,
    metadata || {},
    this.methodInfoGetSomething);
  }

}

