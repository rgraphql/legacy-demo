/* eslint-disable */
import Long from 'long'
import {
  RGQLClientMessage,
  RGQLServerMessage,
} from 'rgraphql'
import _m0 from 'protobufjs/minimal'

export const protobufPackage = 'demo'

/** RPC is the rpc type enum. */
export enum RPC {
  RPC_Unknown = 0,
  RPC_Ping = 1,
  RPC_RGQLClientMessage = 2,
  RPC_RGQLServerMessage = 3,
  UNRECOGNIZED = -1,
}

export function rPCFromJSON(object: any): RPC {
  switch (object) {
    case 0:
    case 'RPC_Unknown':
      return RPC.RPC_Unknown
    case 1:
    case 'RPC_Ping':
      return RPC.RPC_Ping
    case 2:
    case 'RPC_RGQLClientMessage':
      return RPC.RPC_RGQLClientMessage
    case 3:
    case 'RPC_RGQLServerMessage':
      return RPC.RPC_RGQLServerMessage
    case -1:
    case 'UNRECOGNIZED':
    default:
      return RPC.UNRECOGNIZED
  }
}

export function rPCToJSON(object: RPC): string {
  switch (object) {
    case RPC.RPC_Unknown:
      return 'RPC_Unknown'
    case RPC.RPC_Ping:
      return 'RPC_Ping'
    case RPC.RPC_RGQLClientMessage:
      return 'RPC_RGQLClientMessage'
    case RPC.RPC_RGQLServerMessage:
      return 'RPC_RGQLServerMessage'
    case RPC.UNRECOGNIZED:
    default:
      return 'UNRECOGNIZED'
  }
}

/** RPCMessage contains a RPC over the message bus. */
export interface RPCMessage {
  /** RPCId is the RPC identifier. */
  rpcId: RPC
  /** RGQLClientMessage is a client message. */
  rgqlClientMessage: RGQLClientMessage | undefined
  /** RGQLServerMessage is a server message. */
  rgqlServerMessage: RGQLServerMessage | undefined
}

function createBaseRPCMessage(): RPCMessage {
  return {
    rpcId: 0,
    rgqlClientMessage: undefined,
    rgqlServerMessage: undefined,
  }
}

export const RPCMessage = {
  encode(
    message: RPCMessage,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.rpcId !== 0) {
      writer.uint32(8).int32(message.rpcId)
    }
    if (message.rgqlClientMessage !== undefined) {
      RGQLClientMessage.encode(
        message.rgqlClientMessage,
        writer.uint32(18).fork()
      ).ldelim()
    }
    if (message.rgqlServerMessage !== undefined) {
      RGQLServerMessage.encode(
        message.rgqlServerMessage,
        writer.uint32(26).fork()
      ).ldelim()
    }
    return writer
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): RPCMessage {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input)
    let end = length === undefined ? reader.len : reader.pos + length
    const message = createBaseRPCMessage()
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.rpcId = reader.int32() as any
          break
        case 2:
          message.rgqlClientMessage = RGQLClientMessage.decode(
            reader,
            reader.uint32()
          )
          break
        case 3:
          message.rgqlServerMessage = RGQLServerMessage.decode(
            reader,
            reader.uint32()
          )
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  // encodeTransform encodes a source of message objects.
  // Transform<RPCMessage, Uint8Array>
  async *encodeTransform(
    source:
      | AsyncIterable<RPCMessage | RPCMessage[]>
      | Iterable<RPCMessage | RPCMessage[]>
  ): AsyncIterable<Uint8Array> {
    for await (const pkt of source) {
      if (Array.isArray(pkt)) {
        for (const p of pkt) {
          yield* [RPCMessage.encode(p).finish()]
        }
      } else {
        yield* [RPCMessage.encode(pkt).finish()]
      }
    }
  },

  // decodeTransform decodes a source of encoded messages.
  // Transform<Uint8Array, RPCMessage>
  async *decodeTransform(
    source:
      | AsyncIterable<Uint8Array | Uint8Array[]>
      | Iterable<Uint8Array | Uint8Array[]>
  ): AsyncIterable<RPCMessage> {
    for await (const pkt of source) {
      if (Array.isArray(pkt)) {
        for (const p of pkt) {
          yield* [RPCMessage.decode(p)]
        }
      } else {
        yield* [RPCMessage.decode(pkt)]
      }
    }
  },

  fromJSON(object: any): RPCMessage {
    return {
      rpcId: isSet(object.rpcId) ? rPCFromJSON(object.rpcId) : 0,
      rgqlClientMessage: isSet(object.rgqlClientMessage)
        ? RGQLClientMessage.fromJSON(object.rgqlClientMessage)
        : undefined,
      rgqlServerMessage: isSet(object.rgqlServerMessage)
        ? RGQLServerMessage.fromJSON(object.rgqlServerMessage)
        : undefined,
    }
  },

  toJSON(message: RPCMessage): unknown {
    const obj: any = {}
    message.rpcId !== undefined && (obj.rpcId = rPCToJSON(message.rpcId))
    message.rgqlClientMessage !== undefined &&
      (obj.rgqlClientMessage = message.rgqlClientMessage
        ? RGQLClientMessage.toJSON(message.rgqlClientMessage)
        : undefined)
    message.rgqlServerMessage !== undefined &&
      (obj.rgqlServerMessage = message.rgqlServerMessage
        ? RGQLServerMessage.toJSON(message.rgqlServerMessage)
        : undefined)
    return obj
  },

  fromPartial<I extends Exact<DeepPartial<RPCMessage>, I>>(
    object: I
  ): RPCMessage {
    const message = createBaseRPCMessage()
    message.rpcId = object.rpcId ?? 0
    message.rgqlClientMessage =
      object.rgqlClientMessage !== undefined &&
      object.rgqlClientMessage !== null
        ? RGQLClientMessage.fromPartial(object.rgqlClientMessage)
        : undefined
    message.rgqlServerMessage =
      object.rgqlServerMessage !== undefined &&
      object.rgqlServerMessage !== null
        ? RGQLServerMessage.fromPartial(object.rgqlServerMessage)
        : undefined
    return message
  },
}

type Builtin =
  | Date
  | Function
  | Uint8Array
  | string
  | number
  | boolean
  | undefined

export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Long
  ? string | number | Long
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends { $case: string }
  ? { [K in keyof Omit<T, '$case'>]?: DeepPartial<T[K]> } & {
      $case: T['$case']
    }
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>

type KeysOfUnion<T> = T extends T ? keyof T : never
export type Exact<P, I extends P> = P extends Builtin
  ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & Record<
        Exclude<keyof I, KeysOfUnion<P>>,
        never
      >

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any
  _m0.configure()
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined
}
