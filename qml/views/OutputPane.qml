import QtQuick 2.13
import QtQuick.Controls 2.13

import "../."
import "../controls"

Pane {
    id: root
    background: Rectangle {
        color: Style.bgColor
    }
    
    Label {text: mc.output }
}

