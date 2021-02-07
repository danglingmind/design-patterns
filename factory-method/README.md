# Factory Method
Example :
 Logistics app supports different types of modes of logistics eg.
	via truck / ships etc.
	Design should have a factory method which creates a mode-of-transfer it can be
	any.
	All of the returned mode of transfer should always implement the required interface
	which will have deliver method.

NOTES:
	1. There is one creator class(base) with abstract factory method
	2. All the concrete classes will inherit the creator class and will implement their own
	version of factory method
	3. Return type of all the modified factory methods will implement an interface.