import QtQuick 2.13
import QtQuick.Controls 2.13

import "../."
import "../controls"

Pane {

    ListView {
        model: mc.workspaceCtrl.inputCtrl.metadataListModel

        spacing: 10

        width: parent.width
        height: parent.height

        delegate: Row {
            spacing: 10

            property int txtWidth: parent.width / 2 - rmBtn.width / 2 - spacing

            width: parent.width
            height: txtKey.height

            TextField {
                id: txtKey
                width: txtWidth
                placeholderText: "key"
                text: display
            }
            TextField {
                width: txtWidth
                placeholderText: "value"
            }

            CrossButton {
                id: rmBtn
                anchors.verticalCenter: parent.verticalCenter

                color: Style.bgColor3
                rotation: 45
                onClicked: print("clicked")
            }

        }
    }
}
