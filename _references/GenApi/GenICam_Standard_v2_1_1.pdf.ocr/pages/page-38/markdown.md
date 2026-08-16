|  GENICAM |   | ![img-53.jpeg](img-53.jpeg) emva  |
| --- | --- | --- |
|  Version 2.1.1 | Standard  |   |

set to Yes. In the following example the Action node will be invalidated every 10 ms if Action==Active. If the node's value is read while Action==Active the reading will be performed ignoring the cache. As soon as the reading reveals that Action==Idle the polling will stop and the cache will be active again.

<Enumeration Name="Action">
    <EnumEntry Name="Idle">
    <Value>0</Value>
    </EnumEntry>
    <EnumEntry Name="Active">
    <Value>1</Value>
    <IsSelfClearing>Yes</IsSelfClearing>
    </EnumEntry>
    <pValue>ActionReg</pValue>
    <PollingTime>10</PollingTime>
</Enumeration>

#### 2.8.11 StringReg

A string is a (possibly null-terminated) ASCII string placed somewhere in the address space of the camera. A string is exposed via an IString interface. The example below shows how to extract the model name of the camera using a StringReg node. We assume that the ModelName can have a maximum of 128 bytes including the terminating null character.

<StringReg Name="ModelName">
    <Address>0x1234</Address>
    <Length>128</Length>
    <AccessMode>RO</AccessMode>
    <pPort>Device</pPort>
</StringReg>

You can get and set a string through the IString interface.

#### 2.8.12 String (v1.1)

A String node is floating node which can hold any string value.

<String Name="ModelName">
    <Value>This initializes the node's value</Value>
</String>

#### 2.8.13 SwissKnife, IntSwissKnife, Converter, and IntConverter

To do mathematical computations within GenICam, the SwissKnife node dealing with float numbers and the IntSwissKnife node dealing with integers have been introduced. Both have the same syntax.

The following example shows how the product of two numbers is computed. The XTimesY node exposes an IInteger interface reading 504 (=12*42):