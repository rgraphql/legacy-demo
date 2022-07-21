import * as graphql from 'graphql'
import * as rgraphql from 'rgraphql'
import { SoyuzClient } from 'soyuz'

import { RPC, RPCMessage } from './pb/demo.pb'

// DialWebsocketClient dials a SoyuzClient over websocket.
export async function DialWebsocketClient(
  serverURL: string,
  schema: graphql.GraphQLSchema
): Promise<SoyuzClient> {
  const ws = new WebSocket(serverURL)
  return new Promise<SoyuzClient>((resolve, reject) => {
    let client: SoyuzClient
    ws.onopen = (_: Event) => {
      console.log('connected')
      client = new SoyuzClient(schema, (msg: rgraphql.RGQLClientMessage) => {
        // Transmit the message to the server.
        /* tslint:disable-next-line */
        console.log('tx:', msg)
        const data = RPCMessage.encode({
          rpcId: RPC.RPC_RGQLClientMessage,
          rgqlClientMessage: msg,
          rgqlServerMessage: undefined,
        }).finish()
        ws.send(data)
      })
      resolve(client)
    }
    ws.onclose = (_: Event) => {
      console.log('disconnected')
      if (!client) {
        reject(new Error('connection failed'))
      }
    }
    ws.onerror = (err: Event) => {
      if (!client) {
        reject(err)
      }
    }
    ws.onmessage = async (e: MessageEvent) => {
      let data: Uint8Array
      const eventData = e.data
      if (eventData instanceof Uint8Array) {
        data = eventData
      } else {
        const dataBlob: Blob = e.data
        const dataArrayBuffer = await new Response(dataBlob).arrayBuffer()
        data = new Uint8Array(dataArrayBuffer)
      }
      try {
        const msg = RPCMessage.decode(data)
        switch (msg.rpcId) {
          case RPC.RPC_RGQLServerMessage:
            if (msg.rgqlServerMessage) {
              client.handleMessages([msg.rgqlServerMessage])
            }
            break
          case RPC.RPC_Ping:
            break
          default:
            /* tslint:disable-next-line */
            console.error('unhandled rpc type', msg.rpcId)
        }
      } catch (e) {
        /* tslint:disable-next-line */
        console.error('handle message', e)
      }
    }
  })
}
