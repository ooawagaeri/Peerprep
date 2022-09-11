/* eslint-disable */
// @generated by protobuf-ts 2.8.0 with parameter server_grpc1,client_grpc1,eslint_disable,long_type_number
// @generated from protobuf file "collab-service.proto" (package "collaboration_service", syntax proto3)
// tslint:disable
import { ServiceType } from "@protobuf-ts/runtime-rpc";
import type { BinaryWriteOptions } from "@protobuf-ts/runtime";
import type { IBinaryWriter } from "@protobuf-ts/runtime";
import { WireType } from "@protobuf-ts/runtime";
import type { BinaryReadOptions } from "@protobuf-ts/runtime";
import type { IBinaryReader } from "@protobuf-ts/runtime";
import { UnknownFieldHandler } from "@protobuf-ts/runtime";
import type { PartialMessage } from "@protobuf-ts/runtime";
import { reflectionMergePartial } from "@protobuf-ts/runtime";
import { MESSAGE_TYPE } from "@protobuf-ts/runtime";
import { MessageType } from "@protobuf-ts/runtime";
/**
 * @generated from protobuf message collaboration_service.VerifyRoomRequest
 */
export interface VerifyRoomRequest {
    /**
     * @generated from protobuf field: string session_token = 1;
     */
    sessionToken: string;
    /**
     * @generated from protobuf field: string room_token = 2;
     */
    roomToken: string;
}
/**
 * @generated from protobuf message collaboration_service.VerifyRoomResponse
 */
export interface VerifyRoomResponse {
    /**
     * @generated from protobuf field: string error_message = 1;
     */
    errorMessage: string;
    /**
     * @generated from protobuf field: collaboration_service.VerifyRoomErrorCode error_code = 2;
     */
    errorCode: VerifyRoomErrorCode;
}
/**
 * @generated from protobuf enum collaboration_service.VerifyRoomErrorCode
 */
export enum VerifyRoomErrorCode {
    /**
     * @generated from protobuf enum value: VERIFY_ROOM_ERROR_NONE = 0;
     */
    VERIFY_ROOM_ERROR_NONE = 0,
    /**
     * @generated from protobuf enum value: VERIFY_ROOM_BAD_REQUEST = 100;
     */
    VERIFY_ROOM_BAD_REQUEST = 100,
    /**
     * @generated from protobuf enum value: VERIFY_ROOM_INTERNAL_ERROR = 101;
     */
    VERIFY_ROOM_INTERNAL_ERROR = 101,
    /**
     * @generated from protobuf enum value: VERIFY_ROOM_UNAUTHORIZED = 102;
     */
    VERIFY_ROOM_UNAUTHORIZED = 102
}
// @generated message type with reflection information, may provide speed optimized methods
class VerifyRoomRequest$Type extends MessageType<VerifyRoomRequest> {
    constructor() {
        super("collaboration_service.VerifyRoomRequest", [
            { no: 1, name: "session_token", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "room_token", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value?: PartialMessage<VerifyRoomRequest>): VerifyRoomRequest {
        const message = { sessionToken: "", roomToken: "" };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<VerifyRoomRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: VerifyRoomRequest): VerifyRoomRequest {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string session_token */ 1:
                    message.sessionToken = reader.string();
                    break;
                case /* string room_token */ 2:
                    message.roomToken = reader.string();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: VerifyRoomRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string session_token = 1; */
        if (message.sessionToken !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.sessionToken);
        /* string room_token = 2; */
        if (message.roomToken !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.roomToken);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message collaboration_service.VerifyRoomRequest
 */
export const VerifyRoomRequest = new VerifyRoomRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class VerifyRoomResponse$Type extends MessageType<VerifyRoomResponse> {
    constructor() {
        super("collaboration_service.VerifyRoomResponse", [
            { no: 1, name: "error_message", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "error_code", kind: "enum", T: () => ["collaboration_service.VerifyRoomErrorCode", VerifyRoomErrorCode] }
        ]);
    }
    create(value?: PartialMessage<VerifyRoomResponse>): VerifyRoomResponse {
        const message = { errorMessage: "", errorCode: 0 };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<VerifyRoomResponse>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: VerifyRoomResponse): VerifyRoomResponse {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string error_message */ 1:
                    message.errorMessage = reader.string();
                    break;
                case /* collaboration_service.VerifyRoomErrorCode error_code */ 2:
                    message.errorCode = reader.int32();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: VerifyRoomResponse, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string error_message = 1; */
        if (message.errorMessage !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.errorMessage);
        /* collaboration_service.VerifyRoomErrorCode error_code = 2; */
        if (message.errorCode !== 0)
            writer.tag(2, WireType.Varint).int32(message.errorCode);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message collaboration_service.VerifyRoomResponse
 */
export const VerifyRoomResponse = new VerifyRoomResponse$Type();
/**
 * @generated ServiceType for protobuf service collaboration_service.CollabService
 */
export const CollabService = new ServiceType("collaboration_service.CollabService", [
    { name: "VerifyRoom", options: {}, I: VerifyRoomRequest, O: VerifyRoomResponse }
]);
