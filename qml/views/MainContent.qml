import QtQuick 2.13
import QtQuick.Controls 2.13

import "../."

SplitView {
    id: root
    anchors.fill: parent
    orientation: Qt.Horizontal

    handle: Rectangle {
        implicitWidth: 4
        color: Style.borderColor
    }

    // SideBar {
    //     SplitView.minimumWidth: 25
    //     SplitView.preferredWidth: 200
    // }

    InputPane {
        SplitView.fillWidth: true
        SplitView.minimumWidth: root.width / 2
    }

    OutputPane {
        SplitView.minimumWidth: 350
        SplitView.preferredWidth: 500
    }
}

