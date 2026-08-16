![img-10.jpeg](img-10.jpeg)

Table 2-22: Chunk Data Control Summary

|  Name | Level | Interface | Access | Unit | Visibility | Description  |
| --- | --- | --- | --- | --- | --- | --- |
|  ChunkDataControl | R | ICategory | R | - | E | Category that contains the Chunk Data control features.  |
|  ChunkModeActive | R | IBoolean | R/W | - | E | Activates the inclusion of Chunk data in the transmitted payload.  |
|  ChunkXMLEnable | R | IBoolean | R/W | - | E | Activates the inclusion of the GenICam XML necessary to the chunk parser to decode all the Chunk data included in the transmitted payload.  |
|  ChunkSelector | R | IEnumeration | R/W | - | E | Selects which Chunk to enable or control.  |
|  ChunkEnable[ChunkSelector] | R | IBoolean | R/W | - | E | Enables the inclusion of the selected Chunk data in the payload of the image.  |
|  ChunkRegionSelector | O | IEnumeration | R/W | - | E | Selects which Region to retrieve data from.  |
|  ChunkRegionID | O | IEnumeration | R | - | E | Returns the Identifier of Region that the image comes from.  |
|  ChunkRegionIDValue[ChunkRegionSelector] | R | IInteger | R | - | E | Returns the unique integer Identifier value of the Region that the image comes from.  |
|  ChunkComponentSelector | O | IEnumeration | R/W | - | E | Selects the Component from which to retrieve data from.  |
|  ChunkComponentID | O | IEnumeration | R | - | E | Returns the Identifier of the selected Component.  |
|  ChunkComponentIDValue[ChunkComponentSelector] | O | IInteger | R | - | E | Returns a unique Identifier value that corresponds to the selected chunk Component.  |
|  ChunkGroupSelector | O | IEnumeration | R/W | - | E | Selects the component Group from which to retrieve data from.  |
|  ChunkGroupID[ChunkGroupSelector] | O | IEnumeration | R | - | E | Returns a unique Identifier corresponding to the selected Group of components.  |
|  ChunkGroupIDValue[ChunkGroupSelector] | O | IInteger | R | - | E | Returns a unique Identifier value that corresponds to the Group of Components of the selected chunk Component.  |
|  ChunkImageComponent | O | IEnumeration | R | - | I | This feature is deprecated (See ChunkComponentID).  |
|  ChunkPartSelector | O | IInteger | R/W | - | I | This feature is deprecated (See ChunkComponentSelector).  |
|  ChunkImage | R | IRegister | R | - | G | Returns the entire image data included in the payload.  |
|  ChunkOffsetX | R | IInteger | R | - | E | Returns the OffsetX of the image included in the payload.  |