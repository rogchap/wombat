import QtQuick 2.13
import QtQuick.Controls 2.13

import "../."
import "../controls"

Pane {
    id: root

    property variant model: mc.workspaceCtrl.inputCtrl.metadataListModel

    ListView {
        id: lv

        model: root.model

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
                placeholderText: qsTr("key")
                text: display

                onFocusChanged: {
                    if (focus && index + 1 === lv.count) {
                         root.model.addEmpty()
                    }
                }

                onTextChanged: root.model.editKeyAt(index, text)
            }
            TextField {
                width: txtWidth
                placeholderText: qsTr("value")
                text: val

                onTextChanged: root.model.editValAt(index, text) 
            }

            CrossButton {
                id: rmBtn
                anchors.verticalCenter: parent.verticalCenter

                color: Style.bgColor3
                rotation: 45
                onClicked: root.model.removeAt(index)
                visible: lv.count > 1
            }

        }
    }
}
