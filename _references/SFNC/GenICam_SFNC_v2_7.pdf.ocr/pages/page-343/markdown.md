|  Visibility | Beginner  |
| --- | --- |
|  Values | Source0 (If 0 based) Source1 Source2 ... All Device-specific  |

Selects the source to control.

Possible values are:

- Source0: Selects the data source 0.
- Source1: Selects the data source 1.
- Source2: Selects the data source 2.
- ...
- All: Selects all the data sources.

The "All" value can be used to change the features of all the sources at the same time. For example, this can be useful to simultaneously start and stop multiple acquisitions.

### 19.6SourceIDValue

|  Name | SourceIDValue[SourceSelector]  |
| --- | --- |
|  Category | SourceControl  |
|  Level | Optional  |
|  Interface | Integer  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | ≥0  |

Returns a unique Identifier value that correspond to the selected Source.

This value is typically used by the Transport Layer to specify the source of the transmitted data.