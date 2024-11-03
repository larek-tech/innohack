# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: analytics/analytics.proto
# Protobuf Python Version: 5.27.2
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import runtime_version as _runtime_version
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_runtime_version.ValidateProtobufRuntimeVersion(
    _runtime_version.Domain.PUBLIC,
    5,
    27,
    2,
    '',
    'analytics/analytics.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x19\x61nalytics/analytics.proto\x12\tanalytics\"P\n\x06Params\x12\x10\n\x08query_id\x18\x01 \x01(\x03\x12\x12\n\nstart_date\x18\x02 \x01(\t\x12\x10\n\x08\x65nd_date\x18\x03 \x01(\t\x12\x0e\n\x06prompt\x18\x04 \x01(\t\"L\n\x11\x44\x65scriptionReport\x12\x0f\n\x07sources\x18\x01 \x03(\t\x12\x11\n\tfilenames\x18\x02 \x03(\t\x12\x13\n\x0b\x64\x65scription\x18\x03 \x01(\t\"p\n\x0b\x43hartReport\x12 \n\x06\x63harts\x18\x01 \x03(\x0b\x32\x10.analytics.Chart\x12*\n\x0bmultipliers\x18\x02 \x03(\x0b\x32\x15.analytics.Multiplier\x12\x13\n\x0b\x64\x65scription\x18\x03 \x01(\t\"s\n\x05\x43hart\x12\r\n\x05title\x18\x01 \x01(\t\x12\"\n\x04type\x18\x02 \x01(\x0e\x32\x14.analytics.ChartType\x12\x13\n\x0b\x64\x65scription\x18\x03 \x01(\t\x12\"\n\x07records\x18\x04 \x03(\x0b\x32\x11.analytics.Record\"(\n\nMultiplier\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\x01\"\x1e\n\x06Record\x12\t\n\x01x\x18\x01 \x01(\t\x12\t\n\x01y\x18\x02 \x01(\x01*H\n\tChartType\x12\r\n\tUNDEFINED\x10\x00\x12\r\n\tBAR_CHART\x10\x01\x12\r\n\tPIE_CHART\x10\x02\x12\x0e\n\nLINE_CHART\x10\x03\x32\xd2\x01\n\tAnalytics\x12\x38\n\tGetCharts\x12\x11.analytics.Params\x1a\x16.analytics.ChartReport\"\x00\x12K\n\x14GetDescriptionStream\x12\x11.analytics.Params\x1a\x1c.analytics.DescriptionReport\"\x00\x30\x01\x12>\n\x0fGetChartSummary\x12\x11.analytics.Params\x1a\x16.analytics.ChartReport\"\x00\x42\x17Z\x15internal/analytics/pbb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'analytics.analytics_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z\025internal/analytics/pb'
  _globals['_CHARTTYPE']._serialized_start=505
  _globals['_CHARTTYPE']._serialized_end=577
  _globals['_PARAMS']._serialized_start=40
  _globals['_PARAMS']._serialized_end=120
  _globals['_DESCRIPTIONREPORT']._serialized_start=122
  _globals['_DESCRIPTIONREPORT']._serialized_end=198
  _globals['_CHARTREPORT']._serialized_start=200
  _globals['_CHARTREPORT']._serialized_end=312
  _globals['_CHART']._serialized_start=314
  _globals['_CHART']._serialized_end=429
  _globals['_MULTIPLIER']._serialized_start=431
  _globals['_MULTIPLIER']._serialized_end=471
  _globals['_RECORD']._serialized_start=473
  _globals['_RECORD']._serialized_end=503
  _globals['_ANALYTICS']._serialized_start=580
  _globals['_ANALYTICS']._serialized_end=790
# @@protoc_insertion_point(module_scope)
