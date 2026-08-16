|  ChunkSequencerSetActive | R | IInteger | R | - | E | Return the index of the active set of the running sequencer included in the payload.  |
| --- | --- | --- | --- | --- | --- | --- |
|  ChunkScan3dDistanceUnit | O | IEnumeration | R | - | E | Returns the Distance Unit of the payload image.  |
|  ChunkScan3dOutputMode | O | IEnumeration | R | - | E | Returns the Calibrated Mode of the payload image.  |
|  ChunkScan3dCoordinateSystem | O | IEnumeration | R | - | E | Returns the Coordinate System of the image included in the payload.  |
|  ChunkScan3dCoordinateSystemReference | O | IEnumeration | R | - | E | Returns the Coordinate System Position of the image included in the payload.  |
|  ChunkScan3dCoordinateSelector | O | IEnumeration | R/W | - | E | Selects which Coordinate to retrieve data from.  |
|  ChunkScan3dCoordinateScale[ChunkScan3dCoordinateSelector] | O | IFloat | R | - | E | Returns the Scale for the selected coordinate axis of the image included in the payload.  |
|  ChunkScan3dCoordinateOffset[ChunkScan3dCoordinateSelector] | O | IFloat | R | - | E | Returns the Offset for the selected coordinate axis of the image included in the payload.  |
|  ChunkScan3dInvalidDataFlag[ChunkScan3dCoordinateSelector] | O | IBoolean | R | - | E | Returns if a specific non-valid data flag is used in the data stream.  |
|  ChunkScan3dInvalidDataValue[ChunkScan3dCoordinateSelector] | O | IFloat | R | - | E | Returns the Invalid Data Value used for the image included in the payload.  |
|  ChunkScan3dAxisMin[ChunkScan3dCoordinateSelector] | O | IFloat | R | - | E | Returns the Minimum Axis value for the selected coordinate axis of the image included in the payload.  |
|  ChunkScan3dAxisMax[ChunkScan3dCoordinateSelector] | O | IFloat | R | - | E | Returns the Maximum Axis value for the selected coordinate axis of the image included in the payload.  |
|  ChunkScan3dCoordinateTransformSelector | O | IEnumeration | R/W | - | E | Selector for transform values.  |
|  ChunkScan3dTransformValue[ChunkScan3dCoordinateTransformSelector] | O | IFloat | R | - | E | Returns the transform value.  |
|  ChunkScan3dCoordinateReferenceSelector | O | IEnumeration | R/W | - | E | Selector to read a coordinate system reference value defining the transform of a point from one system to the other.  |
|  ChunkScan3dCoordinateReferenceValue[ChunkScan3dCoordinateReferenceSelector] | O | IFloat | R | - | E | Returns the value of a position or pose coordinate for the anchor or transformed coordinate systems relative to the reference point.  |