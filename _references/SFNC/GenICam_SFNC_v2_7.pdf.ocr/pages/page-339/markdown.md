RegionMode[SourceSelector][RegionSelector] = On
RegionDestination[SourceSelector][RegionSelector] = Stream1
Width[SourceSelector][RegionSelector] = 220
Height[SourceSelector][RegionSelector] = 140
TransferSelector = Stream1
TransferControlMode[TransferSelector] = UserControlled
TransferStreamChannel[TransferSelector] = 0

Source 2 Region 2:

SourceSelector = Source2
RegionSelector[SourceSelector] = Region2
RegionMode[SourceSelector][RegionSelector] = On
RegionDestination[SourceSelector][RegionSelector] = Stream2
Width[SourceSelector][RegionSelector] = 220
Height[SourceSelector][RegionSelector] = 330
TransferSelector = Stream2
TransferControlMode[TransferSelector] = UserControlled
TransferStreamChannel[TransferSelector] = 0

Source 1, Region 1 and 2, Transfer and Acquisition control:

TransferSelector = Stream1
TransferStart[TransferSelector]
TransferSelector = Stream2
TransferStart[TransferSelector]
SourceSelector = Source2
AcquisitionStart[SourceSelector]
...
AcquisitionStop[SourceSelector]
TransferSelector = Stream1
TransferStop[TransferSelector]
TransferSelector = Stream2
TransferStop[TransferSelector]