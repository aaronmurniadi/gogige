|  GENICAM |   | ![img-18.jpeg](img-18.jpeg) emva  |
| --- | --- | --- |
|  Version 2.1.1 | Standard  |   |

Some nodes have elements to control accessibility, for example, the register node (see section 2.8.3). In addition, GenICam provides three mechanisms to change the accessibility at runtime:

- A feature can be temporarily locked depending on the value of another node. While locked, a feature is not writable. In terms of the table above, the writable flag is temporarily forced to 0.
- A feature can be temporarily not available depending on the value of another node. In terms of the table above, the writable and the readable flags are temporarily forced to 0.
- A feature can be not implemented at all depending on the value of another node. In terms of the table above, the implemented flag is permanently forced to 0. As oppose to the locked and not available, a feature with not implemented access mode cannot be dynamical changed.

The distinction between being available and being implemented has been made because a GUI might want to handle the two cases differently. A feature being not implemented at all will never be shown to the user and a feature being temporarily not available will be grayed out and the value will be replaced, e.g., by “—”. A temporarily locked feature will be grayed out, but the feature’s value may still be displayed.

A hardware Trigger that can be switched On and Off is a typical example for making a feature temporarily not available. If switched On, an additional feature, the TriggerPolarity, becomes available and denotes whether the hardware signal should be interpreted as an ActiveHigh or an ActiveLow signal. If the Trigger is switched Off, the TriggerPolarity is meaningless and should be grayed out.

Figure 5 shows how this information is handled in the camera description file. The Trigger and the TriggerPolarity feature are implemented using nodes of the Enumeration type that map a set of enumeration entries to integer numbers. For example, the entries for the Trigger feature are On=1 and Off=0. The integer numbers are mapped to registers using nodes of the IntReg type.