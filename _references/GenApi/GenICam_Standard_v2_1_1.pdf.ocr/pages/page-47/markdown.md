|  GENICAM |   | ![img-63.jpeg](img-63.jpeg)emva  |
| --- | --- | --- |
|  Version 2.1.1 | Standard  |   |

#### 2.9.8 ICategory Interface

- NodeList GetFeatures() – returns a list of pointers to the feature nodes

#### 2.9.9 IPort Interface

- void Read( uint8 *pBuffer, int64 Address, int64 Length ) – reads an array of bytes located in the device at [Address, Address+Length]
- void Write( uint8 *pBuffer, int64 Address, int64 Length ) – writes an array of bytes to the device at [Address, Address+Length]

#### 2.9.10 ISelector Interface

- boolean IsSelector() – indicates if that node is a selector
- void GetSelectedFeatures(FeatureList_t &) – returns a list of pointers to the feature nodes which are selected by the current node.
- void GetSelectingFeatures(FeatureList_t &) – returns a list of pointers to the feature nodes which are selecting the current node.