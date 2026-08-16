|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

## Multi-Source Features setting example

Figure 19-1 above presents an example of a relatively complex device supporting multi-source, multi-region with data stream transfer control. This device features two regions of interest for each video source. In this particular example, both regions of interest from Source1 are streamed on the same output stream channel, while a separate stream channel is used for each region of interest of Source2. The three stream channels are transmitted out of the device using the same external Link to the target Host System.

The features setting for this use case is presented below. The detailed description of the feature to control the Sources can be found in section 19.2: "Source Control features below. The features for Region of interest and image format and handling are documented in chapter 4: "Image Format Control" and chapter 20: "Transfer Control" presents the features for explicit control of data Transfer.

So for the particular use case illustrated in Figure 19-1: above, the features would be set to:

Source 1, Region 1:

SourceSelector = Source1
RegionSelector[SourceSelector] = Region1
RegionMode[SourceSelector][RegionSelector] = On
RegionDestination[SourceSelector][RegionSelector] = Stream0
Width[SourceSelector][RegionSelector] = 320
Height[SourceSelector][RegionSelector] = 240

Source 1, Region 2:

SourceSelector = Source1
RegionSelector[SourceSelector] = Region2
RegionMode[SourceSelector][RegionSelector] = On
RegionDestination[SourceSelector][RegionSelector] = Stream0
Width[SourceSelector][RegionSelector] = 420
Height[SourceSelector][RegionSelector] = 340

Source 1, Region 1 and 2, Transfer and Acquisition control:

TransferSelector = Stream0
TransferControlMode[TransferSelector] = UserControlled
TransferStreamChannel[TransferSelector] = 0
TransferStart[TransferSelector]
AcquisitionStart[SourceSelector]
...
AcquisitionStop[SourceSelector]
TransferStop[TransferSelector]

Source 2 Region 1:

SourceSelector = Source2
RegionSelector[SourceSelector] = Region1