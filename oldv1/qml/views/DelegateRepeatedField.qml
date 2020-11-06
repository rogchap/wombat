import QtQuick 2.13
import QtQuick.Controls 2.13

import "../."
import "../controls"

Item {
    id: root

    property alias hintText: hint.text
    property alias delegate: listView.delegate

    height: pane.height + lblRow.height + 14

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
            id: hint
            color: Qt.darker(Style.textColor3, 1.6)
            text: "repeated " + type
        }
        CrossButton {
            text: qsTr("add")
            color: Style.greenColor
            rotation: 45
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

        ListView {
            id: listView

            spacing: 10

            width: parent.width
            height: contentHeight

            model: valueListModel
        }

        Rectangle {
            width: 1
            height: listView.height + 7
            color: Style.accentColor2
            anchors.left: parent.left
            anchors.top: parent.top
            anchors.leftMargin: -7
            anchors.topMargin: -5
        }
    }
}
