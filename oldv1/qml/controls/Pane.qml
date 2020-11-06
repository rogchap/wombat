import QtQuick 2.13
import QtQuick.Controls 2.13

import "../."

Pane {
    id: control

    property alias color: bg.color

    background: Rectangle {
        id: bg
        color: Style.bgColor
    }
}


