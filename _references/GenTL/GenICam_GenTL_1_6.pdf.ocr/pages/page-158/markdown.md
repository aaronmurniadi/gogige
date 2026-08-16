|  ![img-227.jpeg](img-227.jpeg)CAN |   | ![img-228.jpeg](img-228.jpeg)emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

|  Enumerator | Impl | Value | Description  |
| --- | --- | --- | --- |
|   |  |  | return GC_ERR_NOT_AVAILABLE.Data type: UINT64  |
|  BUFFER_PART_INFO_DELIVERED_IMAGEHEIGHT | CM | 11 | The number of lines in the current buffer part as delivered by the transport mechanism. For area scan type images this is usually the number of lines configured in the device. For variable size linescan images this number may be lower than the configured image height. This information refers for example to the information provided in the GigE Vision image stream data trailer. For other technologies, this is to be implemented accordingly.Data type: SIZET  |
|  BUFFER_PART_INFO_REGION_ID | O | 12 | Identifier allowing to group data parts belonging to the same region (usually corresponding with the RegionIDValue/ChunkRegionIDValue features from SFNC). For image based data, all data parts tagged with the same region ID must by definition carry the same region offset/size parameters.It is not mandatory that region ID's within a given buffer make a contiguous sequence of numbers starting with zero.This information refers for example to the information provided in the GigE Vision image stream multi-part data leader.When given part is not tagged with a specific region ID, the function should return GC_ERR_NOT_AVAILABLE.Data type: UINT64  |
|  BUFFER_PART_INFO_DATA_PURPOSE_ID | CM | 13 | Identifier used to tag data parts having the same purpose (usually corresponding with the  |