|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

|  Category | CoaXPress  |
| --- | --- |
|  Level | Recommended  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | ConnectionLockLoss Encoding StreamDataPacketCrc ControlPacketCrc EventPacketCrc DuplicatedCharactersCorrected DuplicatedCharactersUncorrected  |

Selects which Cxp Error Counter to read or reset.

Possible values are:

- ConnectionLockLoss: Counts the number of times the lock was lost.
- Encoding: Counts the number of protocol encoding errors detected.
- StreamDataPacketCrc: Counts the number of CRC errors detected in a data packet. This counter is only available on the CoaXpress host.
- ControlPacketCrc: Counts the number of CRC errors detected in a control packet.
- EventPacketCrc: Counts the number of CRC errors detected in an event packet.
- DuplicatedCharactersCorrected: Counts the number of corrected errors in the duplicated characters in CXP control words.
- DuplicatedCharactersUncorrected: Counts the number of uncorrected errors in the duplicated characters in CXP control words.

### 27.7.20 CxpErrorCounterReset

|  Name | CxpErrorCounterReset[CxpConnectionSelector][CxpErrorCounterSelector]  |
| --- | --- |
|  Category | CoaXPress  |
|  Level | Recommended  |