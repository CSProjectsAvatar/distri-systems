# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: middleware.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x10middleware.proto\x12\x02pb\"\x08\n\x06IpsReq\"\x16\n\x07IPsResp\x12\x0b\n\x03ips\x18\x01 \x03(\t\"\x1a\n\x08StatsReq\x12\x0e\n\x06tourId\x18\x01 \x01(\t\"\xb5\x01\n\tStatsResp\x12\x0f\n\x07matches\x18\x01 \x01(\r\x12/\n\tvictories\x18\x02 \x03(\x0b\x32\x1c.pb.StatsResp.VictoriesEntry\x12\x12\n\nbestPlayer\x18\x03 \x01(\t\x12\x0e\n\x06winner\x18\x04 \x01(\t\x12\x10\n\x08tourName\x18\x05 \x01(\t\x1a\x30\n\x0eVictoriesEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\r:\x02\x38\x01\"3\n\x04\x46ile\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x0c\n\x04\x64\x61ta\x18\x02 \x01(\x0c\x12\x0f\n\x07is_game\x18\x03 \x01(\x08\"X\n\rTournamentReq\x12\x0c\n\x04name\x18\x01 \x01(\t\x12 \n\ttour_type\x18\x02 \x01(\x0e\x32\r.pb.TournType\x12\x17\n\x05\x66iles\x18\x03 \x03(\x0b\x32\x08.pb.File\" \n\x0eTournamentResp\x12\x0e\n\x06tourId\x18\x01 \x01(\t\"\x0b\n\tAllIdsReq\"\x1d\n\nAllIdsResp\x12\x0f\n\x07tourIds\x18\x01 \x03(\t*3\n\x05State\x12\x0f\n\x0bNOT_STARTED\x10\x00\x12\x0b\n\x07STARTED\x10\x01\x12\x0c\n\x08\x46INISHED\x10\x02*8\n\x06Result\x12\x0b\n\x07NOT_RUN\x10\x00\x12\x0b\n\x07P1_WINS\x10\x01\x12\x0b\n\x07P2_WINS\x10\x02\x12\x07\n\x03TIE\x10\x03*9\n\tTournType\x12\x10\n\x0c\x46irst_Defeat\x10\x00\x12\x0e\n\nAll_vs_All\x10\x01\x12\n\n\x06Groups\x10\x02\x32\xf5\x01\n\nMiddleware\x12;\n\x10UploadTournament\x12\x11.pb.TournamentReq\x1a\x12.pb.TournamentResp\"\x00\x12)\n\x08GetStats\x12\x0c.pb.StatsReq\x1a\r.pb.StatsResp\"\x00\x12,\n\tGetAllIds\x12\r.pb.AllIdsReq\x1a\x0e.pb.AllIdsResp\"\x00\x12,\n\x0bGetRndStats\x12\x0c.pb.StatsReq\x1a\r.pb.StatsResp\"\x00\x12#\n\x06GetIPs\x12\n.pb.IpsReq\x1a\x0b.pb.IPsResp\"\x00\x42\x0bZ\t../pb_midb\x06proto3')

_STATE = DESCRIPTOR.enum_types_by_name['State']
State = enum_type_wrapper.EnumTypeWrapper(_STATE)
_RESULT = DESCRIPTOR.enum_types_by_name['Result']
Result = enum_type_wrapper.EnumTypeWrapper(_RESULT)
_TOURNTYPE = DESCRIPTOR.enum_types_by_name['TournType']
TournType = enum_type_wrapper.EnumTypeWrapper(_TOURNTYPE)
NOT_STARTED = 0
STARTED = 1
FINISHED = 2
NOT_RUN = 0
P1_WINS = 1
P2_WINS = 2
TIE = 3
First_Defeat = 0
All_vs_All = 1
Groups = 2


_IPSREQ = DESCRIPTOR.message_types_by_name['IpsReq']
_IPSRESP = DESCRIPTOR.message_types_by_name['IPsResp']
_STATSREQ = DESCRIPTOR.message_types_by_name['StatsReq']
_STATSRESP = DESCRIPTOR.message_types_by_name['StatsResp']
_STATSRESP_VICTORIESENTRY = _STATSRESP.nested_types_by_name['VictoriesEntry']
_FILE = DESCRIPTOR.message_types_by_name['File']
_TOURNAMENTREQ = DESCRIPTOR.message_types_by_name['TournamentReq']
_TOURNAMENTRESP = DESCRIPTOR.message_types_by_name['TournamentResp']
_ALLIDSREQ = DESCRIPTOR.message_types_by_name['AllIdsReq']
_ALLIDSRESP = DESCRIPTOR.message_types_by_name['AllIdsResp']
IpsReq = _reflection.GeneratedProtocolMessageType('IpsReq', (_message.Message,), {
  'DESCRIPTOR' : _IPSREQ,
  '__module__' : 'middleware_pb2'
  # @@protoc_insertion_point(class_scope:pb.IpsReq)
  })
_sym_db.RegisterMessage(IpsReq)

IPsResp = _reflection.GeneratedProtocolMessageType('IPsResp', (_message.Message,), {
  'DESCRIPTOR' : _IPSRESP,
  '__module__' : 'middleware_pb2'
  # @@protoc_insertion_point(class_scope:pb.IPsResp)
  })
_sym_db.RegisterMessage(IPsResp)

StatsReq = _reflection.GeneratedProtocolMessageType('StatsReq', (_message.Message,), {
  'DESCRIPTOR' : _STATSREQ,
  '__module__' : 'middleware_pb2'
  # @@protoc_insertion_point(class_scope:pb.StatsReq)
  })
_sym_db.RegisterMessage(StatsReq)

StatsResp = _reflection.GeneratedProtocolMessageType('StatsResp', (_message.Message,), {

  'VictoriesEntry' : _reflection.GeneratedProtocolMessageType('VictoriesEntry', (_message.Message,), {
    'DESCRIPTOR' : _STATSRESP_VICTORIESENTRY,
    '__module__' : 'middleware_pb2'
    # @@protoc_insertion_point(class_scope:pb.StatsResp.VictoriesEntry)
    })
  ,
  'DESCRIPTOR' : _STATSRESP,
  '__module__' : 'middleware_pb2'
  # @@protoc_insertion_point(class_scope:pb.StatsResp)
  })
_sym_db.RegisterMessage(StatsResp)
_sym_db.RegisterMessage(StatsResp.VictoriesEntry)

File = _reflection.GeneratedProtocolMessageType('File', (_message.Message,), {
  'DESCRIPTOR' : _FILE,
  '__module__' : 'middleware_pb2'
  # @@protoc_insertion_point(class_scope:pb.File)
  })
_sym_db.RegisterMessage(File)

TournamentReq = _reflection.GeneratedProtocolMessageType('TournamentReq', (_message.Message,), {
  'DESCRIPTOR' : _TOURNAMENTREQ,
  '__module__' : 'middleware_pb2'
  # @@protoc_insertion_point(class_scope:pb.TournamentReq)
  })
_sym_db.RegisterMessage(TournamentReq)

TournamentResp = _reflection.GeneratedProtocolMessageType('TournamentResp', (_message.Message,), {
  'DESCRIPTOR' : _TOURNAMENTRESP,
  '__module__' : 'middleware_pb2'
  # @@protoc_insertion_point(class_scope:pb.TournamentResp)
  })
_sym_db.RegisterMessage(TournamentResp)

AllIdsReq = _reflection.GeneratedProtocolMessageType('AllIdsReq', (_message.Message,), {
  'DESCRIPTOR' : _ALLIDSREQ,
  '__module__' : 'middleware_pb2'
  # @@protoc_insertion_point(class_scope:pb.AllIdsReq)
  })
_sym_db.RegisterMessage(AllIdsReq)

AllIdsResp = _reflection.GeneratedProtocolMessageType('AllIdsResp', (_message.Message,), {
  'DESCRIPTOR' : _ALLIDSRESP,
  '__module__' : 'middleware_pb2'
  # @@protoc_insertion_point(class_scope:pb.AllIdsResp)
  })
_sym_db.RegisterMessage(AllIdsResp)

_MIDDLEWARE = DESCRIPTOR.services_by_name['Middleware']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z\t../pb_mid'
  _STATSRESP_VICTORIESENTRY._options = None
  _STATSRESP_VICTORIESENTRY._serialized_options = b'8\001'
  _STATE._serialized_start=491
  _STATE._serialized_end=542
  _RESULT._serialized_start=544
  _RESULT._serialized_end=600
  _TOURNTYPE._serialized_start=602
  _TOURNTYPE._serialized_end=659
  _IPSREQ._serialized_start=24
  _IPSREQ._serialized_end=32
  _IPSRESP._serialized_start=34
  _IPSRESP._serialized_end=56
  _STATSREQ._serialized_start=58
  _STATSREQ._serialized_end=84
  _STATSRESP._serialized_start=87
  _STATSRESP._serialized_end=268
  _STATSRESP_VICTORIESENTRY._serialized_start=220
  _STATSRESP_VICTORIESENTRY._serialized_end=268
  _FILE._serialized_start=270
  _FILE._serialized_end=321
  _TOURNAMENTREQ._serialized_start=323
  _TOURNAMENTREQ._serialized_end=411
  _TOURNAMENTRESP._serialized_start=413
  _TOURNAMENTRESP._serialized_end=445
  _ALLIDSREQ._serialized_start=447
  _ALLIDSREQ._serialized_end=458
  _ALLIDSRESP._serialized_start=460
  _ALLIDSRESP._serialized_end=489
  _MIDDLEWARE._serialized_start=662
  _MIDDLEWARE._serialized_end=907
# @@protoc_insertion_point(module_scope)
