import QtQuick 2.13
import QtQuick.Controls 2.13
import Qt.labs.platform 1.1

MenuBar {
    Menu {
        MenuItem {
            text: qsTr("About")
            onTriggered: about.show()
        }
    }

    About {
        id: about
    }
}

