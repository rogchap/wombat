import QtQuick 2.13
import QtQuick.Controls 2.13

import "../."
import "../controls"

Item {
    id: root

    property string rType: type
    property string rLabel: label

    height: pane.height + lblRow.height + 10

    Row {
        id: lblRow
        height: lbl.height

        spacing: 10 
        anchors.left: parent.left
        anchors.leftMargin: 5

        Label {
            id: lbl
            text: label
        }

        Label { 
            color: Qt.darker(Style.textColor3, 1.6)
            text: "repeated " + type
        }
        CrossButton {
            text: qsTr("add")
            color: Style.greenColor
            onClicked: {
                valueListModel.addValue()
            }
        }
    }

    Pane {
        id: pane
        anchors.top: lblRow.bottom

        width: parent.width
        height: listView.height

        visible: valueListModel.count > 0 
        Component.onCompleted: print(valueListModel.count) 

        ListView {
            id: listView

            spacing: 10

            width: parent.width
            height: contentHeight

            model: valueListModel
            delegate: Item{
                height: dTextField.height
                width: parent.width
                DelegateTextField {
                    id: dTextField
                    hintText: root.rType
                }
                CrossButton {
                    anchors.left: dTextField.right
                    anchors.leftMargin: 5
                    anchors.verticalCenter: parent.verticalCenter

                    color: Style.bgColor3
                    rotation: 45

                    onClicked: valueListModel.remove(index)
                }
            }
        }

        Rectangle {
            width: 1
            height: listView.height + 5
            color: Style.accentColor2
            anchors.left: parent.left
            anchors.top: parent.top
            anchors.leftMargin: -7
            anchors.topMargin: -5
        }
    }
}
