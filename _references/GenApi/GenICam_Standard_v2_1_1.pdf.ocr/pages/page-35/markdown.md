|  GENICAM |   | ![img-50.jpeg](img-50.jpeg) emva  |
| --- | --- | --- |
|  Version 2.1.1 | Standard  |   |

<Boolean Name="Trigger">
    TriggerReg

    <OnValue>1</OnValue>
    <OffValue>0</OffValue>
</Boolean>

<IntReg Name="TriggerReg">
    <Address>0x6789</Address>
    <Length>1</Length>
    <AccessMode>RW</AccessMode>
    Device
    <Sign>Unsigned</Sign>
    <Endianess>BigEndian</Endianess>
</IntReg>

The Boolean node's value is either taken from another node referenced by a <pValue> entry or holds its own value initialized by the content of the <Value> entry.</Value>

#### 2.8.8 Command

The ICommand interface lets the user submit a command by calling the method Execute and then poll to learn if the execution has been accomplished by calling the method IsDone. Note that IsDone always read the device register regardless of the state of the caches.

The corresponding Command node inherits the elements and attributes of the Node node.

In addition it has a CommandValue element which holds an integer constant which is written into a node which is referenced to by a pValue element. Writing the command value submits the command. IsDone reads the value back and returns false as long as return value equals the command value. If the node is WriteOnly IsDone always returns true. In order to make a floating Command node possible instead of a pValue element also a Value element is allowed. The CommandValue can alternatively also be taken from pCommandValue. A <PollingTime> entry can be used to handle self-clearing commands: While the command is active the node is invalidated each time the PollingTime expires. If a call to IsDone reveals that the command is gone idle the polling stops. If the command is write only no polling takes place. The state of other nodes can depend on the completion of the Command therefore applications that do not implement the event based polling should always wait for the command completion by calling IsDone immediately after the Execute.</PollingTime>

#### 2.8.9 Float, FloatReg

The IFloat interface has a definition similar to the definition of the IInteger interface as described in the section above. It has a Value that is restricted by the Minimum and Maximum parameters, but in contrast to integer, the increment exists only optional and is not verified an writing. Note that the increment does only make sense if it is a constant. In addition, IFloat exposes a Unit that is just a string for display purposes.

The Float node is built analogously to the Integer node in that it has the <Value>, <Min>, <Max>, <Inc>, or <pValue>, <pMin>, <pMax>, <pInc> restriction parameters respectively. In addition, it can have a <Representation> element that can take the values Linear, Logarithmic, or PureNumber, a <Unit> element that contains the unit as a string, a <DisplayNotation> element which can have the value Automatic, Fixed, and Scientific, and a</DisplayNotation></Unit></pMax></pMin></pMax></pMin></Max></pMin></pMax></pMin></pValue>