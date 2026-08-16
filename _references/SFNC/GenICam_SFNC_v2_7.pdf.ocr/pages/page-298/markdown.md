Here below, in addition to EventControl, EventSelector and EventNotification should be listed all the categories and data related features for each Event listed in the EventSelector enumeration feature.

For simplicity, all the categories and their data members are not listed explicitly in that document but a precise naming convention for the categories and their member is provided above instead.

Below, the detailed features for the members of the EventSelector are only listed for 4 typically recommended events: FrameTrigger, ExposureEnd, Error and EventTest.

All the other members of the EventSelector feature should follow the exact same pattern for their features naming and category if they are present in a device.

### 15.1 EventControl

|  Name | EventControl  |
| --- | --- |
|  Category | Root  |
|  Level | Recommended  |
|  Interface | ICategory  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | -  |

Category that contains Event control features.

### 15.2 EventSelector

|  Name | EventSelector  |
| --- | --- |
|  Category | EventControl  |
|  Level | Recommended  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |