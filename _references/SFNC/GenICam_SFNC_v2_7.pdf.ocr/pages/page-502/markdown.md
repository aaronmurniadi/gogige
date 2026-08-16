|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

Note: Device is the name of the standard port that is used to connect the node map to the transport layer and access the control port of the device. Device is a port node (not a feature node) and is generally not accessed by the end user directly. Device must not be included in the root feature tree.

### 26.1.1 ValueArrayCandidates

|  Name | ValueArrayCandidates  |
| --- | --- |
|  Category | None  |
|  Level | Optional  |
|  Interface | IString  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Invisible  |
|  Values | -  |

Advertises selector/value feature pairs suitable for use with GenApi ValueArrayAdapter utility which allows efficient extraction of features under selectors (for example Chunk features).

The features are listed in a comma separated list with the syntax FeatureA[SelectorB], for instance ChunkTimestamp[ChunkScanLineSelector].

The string is expected to be static, not changing its value depending on other features.

The ValueArrayAdapter utility provides optimized access to a (large) array of values that has to be otherwise read using the value feature itself and associated selector (acting as the array index). If the array is large and reasonably compact (non-sparse), the adapter can bring significant performance advantage.