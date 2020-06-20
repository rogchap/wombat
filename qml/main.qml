import QtQuick 2.13
import QtQuick.Controls 2.13

ApplicationWindow {
    id: window
    visible: true
    title: "Courier"
    minimumWidth: 800
    minimumHeight: 450

    Shortcut {
        sequence: "r"
        onActivated: {
            window.close();
            eCtx.reload();
        }
    }

    SplitView {
        anchors.fill: parent
        orientation: Qt.Horizontal
        handle: Rectangle {
            id: handle
            implicitWidth: 1
        }

        Rectangle {
            implicitWidth: 200
            SplitView.maximumWidth: 400
            color: "lightblue"
            Label {
                text: "View 1"
                anchors.centerIn: parent
            }
        }
        Rectangle {
            id: centerItem
            SplitView.minimumWidth: 50
            SplitView.fillWidth: true
            color: "lightgray"
            Label {
                text: "View 2"
                anchors.centerIn: parent
            }
        }
        Rectangle {
            implicitWidth: 200
            color: "lightgreen"
            Label {
                text: "View 3"
                anchors.centerIn: parent
            }
        }
    }
}
