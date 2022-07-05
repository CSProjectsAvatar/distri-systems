"""
@generated by mypy-protobuf.  Do not edit manually!
isort:skip_file
"""
import builtins
import google.protobuf.descriptor
import google.protobuf.internal.containers
import google.protobuf.internal.enum_type_wrapper
import google.protobuf.message
import typing
import typing_extensions

DESCRIPTOR: google.protobuf.descriptor.FileDescriptor

class _State:
    ValueType = typing.NewType('ValueType', builtins.int)
    V: typing_extensions.TypeAlias = ValueType
class _StateEnumTypeWrapper(google.protobuf.internal.enum_type_wrapper._EnumTypeWrapper[_State.ValueType], builtins.type):
    DESCRIPTOR: google.protobuf.descriptor.EnumDescriptor
    NOT_STARTED: _State.ValueType  # 0
    STARTED: _State.ValueType  # 1
    FINISHED: _State.ValueType  # 2
class State(_State, metaclass=_StateEnumTypeWrapper):
    """RunTournament
    message RunReq {
       string name = 1;
    }

    message RunResp {
       repeated Match matchs = 1;
    }
    message Match {
       repeated string players = 1;
       State state = 2;
       Result result = 3;
    }

    """
    pass

NOT_STARTED: State.ValueType  # 0
STARTED: State.ValueType  # 1
FINISHED: State.ValueType  # 2
global___State = State


class _Result:
    ValueType = typing.NewType('ValueType', builtins.int)
    V: typing_extensions.TypeAlias = ValueType
class _ResultEnumTypeWrapper(google.protobuf.internal.enum_type_wrapper._EnumTypeWrapper[_Result.ValueType], builtins.type):
    DESCRIPTOR: google.protobuf.descriptor.EnumDescriptor
    NOT_RUN: _Result.ValueType  # 0
    P1_WINS: _Result.ValueType  # 1
    P2_WINS: _Result.ValueType  # 2
    TIE: _Result.ValueType  # 3
class Result(_Result, metaclass=_ResultEnumTypeWrapper):
    pass

NOT_RUN: Result.ValueType  # 0
P1_WINS: Result.ValueType  # 1
P2_WINS: Result.ValueType  # 2
TIE: Result.ValueType  # 3
global___Result = Result


class _TournType:
    ValueType = typing.NewType('ValueType', builtins.int)
    V: typing_extensions.TypeAlias = ValueType
class _TournTypeEnumTypeWrapper(google.protobuf.internal.enum_type_wrapper._EnumTypeWrapper[_TournType.ValueType], builtins.type):
    DESCRIPTOR: google.protobuf.descriptor.EnumDescriptor
    First_Defeat: _TournType.ValueType  # 0
    All_vs_All: _TournType.ValueType  # 1
class TournType(_TournType, metaclass=_TournTypeEnumTypeWrapper):
    pass

First_Defeat: TournType.ValueType  # 0
All_vs_All: TournType.ValueType  # 1
global___TournType = TournType


class StatsReq(google.protobuf.message.Message):
    """GetStats"""
    DESCRIPTOR: google.protobuf.descriptor.Descriptor
    TOURID_FIELD_NUMBER: builtins.int
    tourId: typing.Text
    def __init__(self,
        *,
        tourId: typing.Text = ...,
        ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["tourId",b"tourId"]) -> None: ...
global___StatsReq = StatsReq

class StatsResp(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor
    class VictoriesEntry(google.protobuf.message.Message):
        DESCRIPTOR: google.protobuf.descriptor.Descriptor
        KEY_FIELD_NUMBER: builtins.int
        VALUE_FIELD_NUMBER: builtins.int
        key: typing.Text
        value: builtins.int
        def __init__(self,
            *,
            key: typing.Text = ...,
            value: builtins.int = ...,
            ) -> None: ...
        def ClearField(self, field_name: typing_extensions.Literal["key",b"key","value",b"value"]) -> None: ...

    MATCHES_FIELD_NUMBER: builtins.int
    VICTORIES_FIELD_NUMBER: builtins.int
    BESTPLAYER_FIELD_NUMBER: builtins.int
    WINNER_FIELD_NUMBER: builtins.int
    TOURNAME_FIELD_NUMBER: builtins.int
    matches: builtins.int
    @property
    def victories(self) -> google.protobuf.internal.containers.ScalarMap[typing.Text, builtins.int]: ...
    bestPlayer: typing.Text
    winner: typing.Text
    tourName: typing.Text
    def __init__(self,
        *,
        matches: builtins.int = ...,
        victories: typing.Optional[typing.Mapping[typing.Text, builtins.int]] = ...,
        bestPlayer: typing.Text = ...,
        winner: typing.Text = ...,
        tourName: typing.Text = ...,
        ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["bestPlayer",b"bestPlayer","matches",b"matches","tourName",b"tourName","victories",b"victories","winner",b"winner"]) -> None: ...
global___StatsResp = StatsResp

class File(google.protobuf.message.Message):
    """UploadTournament"""
    DESCRIPTOR: google.protobuf.descriptor.Descriptor
    NAME_FIELD_NUMBER: builtins.int
    DATA_FIELD_NUMBER: builtins.int
    IS_GAME_FIELD_NUMBER: builtins.int
    name: typing.Text
    data: builtins.bytes
    is_game: builtins.bool
    def __init__(self,
        *,
        name: typing.Text = ...,
        data: builtins.bytes = ...,
        is_game: builtins.bool = ...,
        ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["data",b"data","is_game",b"is_game","name",b"name"]) -> None: ...
global___File = File

class TournamentReq(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor
    NAME_FIELD_NUMBER: builtins.int
    TOUR_TYPE_FIELD_NUMBER: builtins.int
    FILES_FIELD_NUMBER: builtins.int
    name: typing.Text
    tour_type: global___TournType.ValueType
    @property
    def files(self) -> google.protobuf.internal.containers.RepeatedCompositeFieldContainer[global___File]: ...
    def __init__(self,
        *,
        name: typing.Text = ...,
        tour_type: global___TournType.ValueType = ...,
        files: typing.Optional[typing.Iterable[global___File]] = ...,
        ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["files",b"files","name",b"name","tour_type",b"tour_type"]) -> None: ...
global___TournamentReq = TournamentReq

class TournamentResp(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor
    TOURID_FIELD_NUMBER: builtins.int
    tourId: typing.Text
    def __init__(self,
        *,
        tourId: typing.Text = ...,
        ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["tourId",b"tourId"]) -> None: ...
global___TournamentResp = TournamentResp

class AllIdsReq(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor
    def __init__(self,
        ) -> None: ...
global___AllIdsReq = AllIdsReq

class AllIdsResp(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor
    TOURIDS_FIELD_NUMBER: builtins.int
    @property
    def tourIds(self) -> google.protobuf.internal.containers.RepeatedScalarFieldContainer[typing.Text]: ...
    def __init__(self,
        *,
        tourIds: typing.Optional[typing.Iterable[typing.Text]] = ...,
        ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["tourIds",b"tourIds"]) -> None: ...
global___AllIdsResp = AllIdsResp
