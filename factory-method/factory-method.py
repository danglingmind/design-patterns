import abc

# this can be set using business logics
DefaultMode = 'ship'


class DiliveryMode:
    def __init__(self):
        pass

    @abc.abstractmethod
    def Diliver(self):
        pass


class TruckLogistics(DiliveryMode):
    def __init__(self):
        pass

    def Diliver(self):
        print("Dilivering via Truck")


class ShipLogistics(DiliveryMode):
    def __init__(self):
        pass

    def Diliver(self):
        print("Dilivering via Ship")


class Application:
    def __init__(self):
        pass

    def DiliveryAgent(self):
        if DefaultMode == 'truck':  # this could be set using any other business logics
            return TruckLogistics()
        else:
            return ShipLogistics()


if __name__ == "__main__":
    myApp = Application()

    diliveryAgent = myApp.DiliveryAgent()

    diliveryAgent.Diliver()
