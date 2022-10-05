const enum ConnectionOpCode {
  DATA,
  JOIN,
  ACK,
  ARP,
  WMI,
}

declare type TunnelMessage = {
  sender: string,
  data: Uint8Array,
  flag: ConnectionOpCode,
};

declare type TunnelInternalMessage = {
  sender: string,
  data: Array<number>,
  flag: ConnectionOpCode,
};

export {
  ConnectionOpCode,
  TunnelMessage,
  TunnelInternalMessage,
};
