|  GENICAM |   | ![img-57.jpeg](img-57.jpeg) emva  |
| --- | --- | --- |
|  Version 2.1.1 | Standard  |   |

<Category Name="Root">
    <pFeature>CommandRegBase</pFeature>
    <pFeature>VendorName</pFeature>
    <pFeature>ModelName</pFeature>
</Category>

<ConfRom Name="ConfRom">
    <Unit>0x00A02D</Unit>
    <Address>0x400</Address>
    <pAddress>InitialNodeSpace</pAddress>
    <Length>0x400</Length>
    <pPort>Device</pPort>
    <IntKey Name="CommandRegBase">0x40</IntKey>
    <TextDesc Name="VendorName">0x81</TextDesc>
    <TextDesc Name="ModelName">0x82</TextDesc>
</ConfRom>

<Integer Name="InitialNodeSpace">
    <Value>0xFFFFF0000000</Value>
</Integer>

Note that a ConfROM node has <address>, <pAddress>, <IntSwissKnife>, <Length>, and <pPort> elements that have the same meaning as with other Registers (see section 2.8.3). </pPort></IntSwissKnife></pAddress></Address>

The typical implementation will create separate nodes for the <intkey> and the <textdesc> elements that are given the name denoted in the respective entry's Name attribute, a <p1212parser> element pointing to the ConfROM node and a <key> element with the respective key values.</key></textdesc></intkey>

#### 2.8.15 DcamLock and SmartFeature

Currently, most standard register layouts are fixed mechanisms, and methods are required to give access to custom features not defined in the standard. GenICam currently supports two access mechanisms.

The DcamLock node can retrieve the address of a smart feature exposed according to the DCAM advanced features mechanism. It inherits the elements and attributes from the Register node. The following example unlocks an advanced DCAM feature with a <FeatureID> element of 0x0030533B73C3 where 0x003053 is a vendor ID and 0x3B73C3 is a feature ID defined by that vendor. The value 0 in the <Timeout> element means that the feature will not unlock automatically.</Timeout>

<advfeaturelock name="BaslerAdvFeatureLock">
    <FeatureID>0x0030533B73C3</FeatureID>
    <Timeout>0</Timeout>
    <Address>0xfffff2f00000</Address>
    <Length>8</Length>
    <AccessMode>RW</AccessMode>
    <pPort>Device</pPort>
</advfeaturelock>