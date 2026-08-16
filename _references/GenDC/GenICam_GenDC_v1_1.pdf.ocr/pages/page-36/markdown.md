example putting the data to the memory of a graphics card and keeping the Descriptor in normal Host memory. For this, the internal offsets of the GenDC Data Container do not need to be altered; all is done by changing the Flow base addresses which are not part of the Container.

Early processing of Parts transmitted in different Flows can be easily done if the Transport Layer provides an end of Flow notification of some kind. It is highly recommended for each Transport Layer to provide such an end of Flow notification.

[ R-003] A GenDC compliant transmitter must provide a Flow mapping table (stored in little-endian ordering).