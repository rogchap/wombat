import QtQuick 2.13
import QtQuick.Controls 2.13

import "../."
import "."

Item {
    id: control

    property alias labelText: label.text
    property alias model: combobox.model
    property alias displayText: combobox.displayText
    property alias currentIndex: combobox.currentIndex
    property int labelLeftMargin: 5

    height: label.height + combobox.height + 5
    implicitWidth: combobox.width

    Label {
        id: label

        anchors {
            left: control.left
            leftMargin: labelLeftMargin
        }
    }

    ComboBox {
        id: combobox

        anchors {
            top: label.bottom
            topMargin: 5
        }

        background: Rectangle {
            implicitHeight: 40
            implicitWidth: 400
            color: Style.bgInputColor
            border.color: Style.borderColor
        }
    } 
}


