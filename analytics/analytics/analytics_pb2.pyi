from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class ChartType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    UNDEFINED: _ClassVar[ChartType]
    BAR_CHART: _ClassVar[ChartType]
    PIE_CHART: _ClassVar[ChartType]
UNDEFINED: ChartType
BAR_CHART: ChartType
PIE_CHART: ChartType

class Params(_message.Message):
    __slots__ = ("query_id", "start_date", "end_date", "prompt")
    QUERY_ID_FIELD_NUMBER: _ClassVar[int]
    START_DATE_FIELD_NUMBER: _ClassVar[int]
    END_DATE_FIELD_NUMBER: _ClassVar[int]
    PROMPT_FIELD_NUMBER: _ClassVar[int]
    query_id: int
    start_date: str
    end_date: str
    prompt: str
    def __init__(self, query_id: _Optional[int] = ..., start_date: _Optional[str] = ..., end_date: _Optional[str] = ..., prompt: _Optional[str] = ...) -> None: ...

class Report(_message.Message):
    __slots__ = ("source", "filename", "description", "charts", "multipliers")
    SOURCE_FIELD_NUMBER: _ClassVar[int]
    FILENAME_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    CHARTS_FIELD_NUMBER: _ClassVar[int]
    MULTIPLIERS_FIELD_NUMBER: _ClassVar[int]
    source: str
    filename: str
    description: str
    charts: _containers.RepeatedCompositeFieldContainer[Chart]
    multipliers: _containers.RepeatedCompositeFieldContainer[Multiplier]
    def __init__(self, source: _Optional[str] = ..., filename: _Optional[str] = ..., description: _Optional[str] = ..., charts: _Optional[_Iterable[_Union[Chart, _Mapping]]] = ..., multipliers: _Optional[_Iterable[_Union[Multiplier, _Mapping]]] = ...) -> None: ...

class Chart(_message.Message):
    __slots__ = ("title", "type", "description", "records")
    TITLE_FIELD_NUMBER: _ClassVar[int]
    TYPE_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    RECORDS_FIELD_NUMBER: _ClassVar[int]
    title: str
    type: ChartType
    description: str
    records: _containers.RepeatedCompositeFieldContainer[Record]
    def __init__(self, title: _Optional[str] = ..., type: _Optional[_Union[ChartType, str]] = ..., description: _Optional[str] = ..., records: _Optional[_Iterable[_Union[Record, _Mapping]]] = ...) -> None: ...

class Multiplier(_message.Message):
    __slots__ = ("key", "value")
    KEY_FIELD_NUMBER: _ClassVar[int]
    VALUE_FIELD_NUMBER: _ClassVar[int]
    key: str
    value: float
    def __init__(self, key: _Optional[str] = ..., value: _Optional[float] = ...) -> None: ...

class Record(_message.Message):
    __slots__ = ("x", "y")
    X_FIELD_NUMBER: _ClassVar[int]
    Y_FIELD_NUMBER: _ClassVar[int]
    x: str
    y: float
    def __init__(self, x: _Optional[str] = ..., y: _Optional[float] = ...) -> None: ...
