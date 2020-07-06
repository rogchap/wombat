import QtQuick 2.13
import QtQuick.Controls 2.13

import "../."

SplitView {
    id: root
    anchors.fill: parent
    orientation: Qt.Horizontal

    handle: Rectangle {
        implicitWidth: 1
        color: Style.borderColor
    }

    SideBar {
        SplitView.minimumWidth: 25
        SplitView.preferredWidth: 200
    }

    InputPane {
        SplitView.fillWidth: true
    }

    OutputPane {
        SplitView.minimumWidth: 25
        SplitView.preferredWidth: 400
    }
}

