|  ![img-225.jpeg](img-225.jpeg)CAN |   | ![img-226.jpeg](img-226.jpeg)emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

|  Enumerator | Impl | Value | Description  |
| --- | --- | --- | --- |
|   |  |  | to handle areas of interest.If the information is not applicable to given data type, the query should result in GC_ERR_NOT_AVAILABLE.Data type: SIZET  |
|  BUFFER_PART_INFO_XPADDING | CM | 9 | XPadding of the data in the buffer part in number of pixels.If the information is not applicable to given data type, the query should result in GC_ERR_NOT_AVAILABLE.Data type: SIZET  |
|  BUFFER_PART_INFO_SOURCE_ID | O | 10 | Identifier allowing to group data parts belonging to the same source (usually corresponding with the SourceIDValue/ChunkSourceIDValue features from SFNC). Parts marked with the same source ID can be pixel-mapped together. Parts carrying data from different ROI's of the same source would typically be marked with the same source ID.It is not mandatory that source ID's within a given buffer make a contiguous sequence of numbers starting with zero.Note: for example with a dual-source 3D camera, the buffer can contain data parts carrying the 3D data and Confidence data corresponding to both of the two different sources. In this case the source ID helps to match the 3D and Confidence parts belonging together.This information refers for example to the information provided in the GigE Vision image stream multi-part data leader.When given part is not tagged with a specific source ID, the function should  |