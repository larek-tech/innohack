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
    LINE_CHART: _ClassVar[ChartType]
UNDEFINED: ChartType
BAR_CHART: ChartType
PIE_CHART: ChartType
LINE_CHART: ChartType

class Params(_message.Message):
    __slots__ = ("query_id", "prompt")
    QUERY_ID_FIELD_NUMBER: _ClassVar[int]
    PROMPT_FIELD_NUMBER: _ClassVar[int]
    query_id: int
    prompt: str
    def __init__(self, query_id: _Optional[int] = ..., prompt: _Optional[str] = ...) -> None: ...

class DescriptionReport(_message.Message):
    __slots__ = ("sources", "filenames", "description")
    SOURCES_FIELD_NUMBER: _ClassVar[int]
    FILENAMES_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    sources: _containers.RepeatedScalarFieldContainer[str]
    filenames: _containers.RepeatedScalarFieldContainer[str]
    description: str
    def __init__(self, sources: _Optional[_Iterable[str]] = ..., filenames: _Optional[_Iterable[str]] = ..., description: _Optional[str] = ...) -> None: ...

class Filter(_message.Message):
    __slots__ = ("start_date", "end_date")
    START_DATE_FIELD_NUMBER: _ClassVar[int]
    END_DATE_FIELD_NUMBER: _ClassVar[int]
    start_date: int
    end_date: int
    def __init__(self, start_date: _Optional[int] = ..., end_date: _Optional[int] = ...) -> None: ...

class ChartReport(_message.Message):
    __slots__ = ("charts", "multipliers", "description")
    CHARTS_FIELD_NUMBER: _ClassVar[int]
    MULTIPLIERS_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    charts: _containers.RepeatedCompositeFieldContainer[Chart]
    multipliers: _containers.RepeatedCompositeFieldContainer[Multiplier]
    description: str
    def __init__(self, charts: _Optional[_Iterable[_Union[Chart, _Mapping]]] = ..., multipliers: _Optional[_Iterable[_Union[Multiplier, _Mapping]]] = ..., description: _Optional[str] = ...) -> None: ...

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
