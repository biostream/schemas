# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: bmeg/protein.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='bmeg/protein.proto',
  package='bmeg',
  syntax='proto3',
  serialized_pb=_b('\n\x12\x62meg/protein.proto\x12\x04\x62meg\"\xa6\x01\n\x07Protein\x12\x12\n\nuniprot_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x11\n\taccession\x18\x03 \x03(\t\x12\x14\n\x0c\x65nsembl_gene\x18\x04 \x01(\t\x12\x1a\n\x12\x65nsembl_transcript\x18\x05 \x01(\t\x12\x17\n\x0f\x65nsembl_protein\x18\x06 \x01(\t\x12\x0e\n\x06pubmed\x18\x07 \x03(\t\x12\x0b\n\x03pdb\x18\x08 \x03(\t\"\x81\x01\n\nPFAMFamily\x12\n\n\x02id\x18\x01 \x01(\t\x12\x11\n\taccession\x18\x02 \x01(\t\x12\x0c\n\x04type\x18\x03 \x01(\t\x12\x10\n\x08go_terms\x18\x04 \x03(\t\x12\r\n\x05\x63lans\x18\x05 \x03(\t\x12\x13\n\x0b\x64\x65scription\x18\x06 \x01(\t\x12\x10\n\x08\x63omments\x18\x07 \x01(\tb\x06proto3')
)




_PROTEIN = _descriptor.Descriptor(
  name='Protein',
  full_name='bmeg.Protein',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='uniprot_id', full_name='bmeg.Protein.uniprot_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='name', full_name='bmeg.Protein.name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='accession', full_name='bmeg.Protein.accession', index=2,
      number=3, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='ensembl_gene', full_name='bmeg.Protein.ensembl_gene', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='ensembl_transcript', full_name='bmeg.Protein.ensembl_transcript', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='ensembl_protein', full_name='bmeg.Protein.ensembl_protein', index=5,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='pubmed', full_name='bmeg.Protein.pubmed', index=6,
      number=7, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='pdb', full_name='bmeg.Protein.pdb', index=7,
      number=8, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=29,
  serialized_end=195,
)


_PFAMFAMILY = _descriptor.Descriptor(
  name='PFAMFamily',
  full_name='bmeg.PFAMFamily',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='bmeg.PFAMFamily.id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='accession', full_name='bmeg.PFAMFamily.accession', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='type', full_name='bmeg.PFAMFamily.type', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='go_terms', full_name='bmeg.PFAMFamily.go_terms', index=3,
      number=4, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='clans', full_name='bmeg.PFAMFamily.clans', index=4,
      number=5, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='description', full_name='bmeg.PFAMFamily.description', index=5,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='comments', full_name='bmeg.PFAMFamily.comments', index=6,
      number=7, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=198,
  serialized_end=327,
)

DESCRIPTOR.message_types_by_name['Protein'] = _PROTEIN
DESCRIPTOR.message_types_by_name['PFAMFamily'] = _PFAMFAMILY
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Protein = _reflection.GeneratedProtocolMessageType('Protein', (_message.Message,), dict(
  DESCRIPTOR = _PROTEIN,
  __module__ = 'bmeg.protein_pb2'
  # @@protoc_insertion_point(class_scope:bmeg.Protein)
  ))
_sym_db.RegisterMessage(Protein)

PFAMFamily = _reflection.GeneratedProtocolMessageType('PFAMFamily', (_message.Message,), dict(
  DESCRIPTOR = _PFAMFAMILY,
  __module__ = 'bmeg.protein_pb2'
  # @@protoc_insertion_point(class_scope:bmeg.PFAMFamily)
  ))
_sym_db.RegisterMessage(PFAMFamily)


# @@protoc_insertion_point(module_scope)
