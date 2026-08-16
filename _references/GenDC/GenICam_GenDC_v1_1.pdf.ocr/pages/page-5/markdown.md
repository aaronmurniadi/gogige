## List of Tables

Table 1-1: Terms and Definitions ... 8
Table 1-2: Requirements terminology... 9
Table 2-1: GenDC Container Header Description... 16
Table 2-2: GenDC Component Header fields description ... 19
Table 2-3: GenDC Component Header Part common fields description... 22
Table 2-4: GenDC Component Header Part type specific fields description... 23
Table 2-5: GenDC Part Types ... 25
Table 2-6: Part specific Header fields description for Metadata ... 26
Table 2-7: Part specific Header fields description for 1D Array or Image ... 29
Table 2-8: Part specific Header fields description for 2D uncompressed, JPEG or JPEG2000 compressed image ... 29
Table 2-9: Part specific Header fields description for H.264 compressed image... 32
Table 3-1: GenDC Flow mapping table description ... 35

## List of Requirements and Objectives

[ R-001] A GenDC compliant product must use the headers and flags as defined by this specification. ... 12
[ R-002] A GenDC compliant product must use the Part header types as defined by this specification. ... 21
[ R-003] A GenDC compliant transmitter must provide a Flow mapping table (stored in little-endian ordering). ... 36
[ R-004] The GenDC Container Descriptor must be always stored in little-endian ordering. ... 37
[ R-005] Container data Part in PFNC format must use little-endian ordering. ... 37
[ R-006] For Transmission and file storage, a GenDC Container is always represented as a continuous block of linear
memory starting with the Descriptor. ... 37
[ R-007] The Part's DataOffset is always the offset of the data in bytes from the start of the Descriptor. ... 37
[ R-008] The Part's FlowOffset is always the offset of the data in bytes from the start of the Flow specified by the Part's
FlowId. ... 37
[ R-009] The Descriptor is always transferred in Flow 0. ... 38
[ R-010] A Part must only be mapped to a single Flow. ... 38
[ R-011] Flow must be numbered sequentially starting from 0. ... 38
[ R-014] When storing a GenDC Container using the ".gendc" file extension, the standard GenDC binary format must be
used. ... 38
[ R-015] When storing a GenDC Container, a linear Container must be used. ... 38
[ R-016] When adding metadata to a Container, this must be done using a separate Metadata Component. ... 38
[ CR-012] If a Container has variable content during the transmission the VariableFields flags of the Container must be
set accordingly. ... 38
[ CR-013] If any of the Container's VariableFields flags is set, a final Descriptor must be sent as soon as possible but at
least just after the transmission of the data section. ... 38