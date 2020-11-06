<script>
	import { createEventDispatcher } from 'svelte';
	const dispatch = createEventDispatcher();
	export let type = "horizontal";
	export let pos = 50;
	export let fixed = false;
	export let min = 400;
	// export let min1 = min;
	// export let min2 = min;
	const refs = {};
	let dragging = false;

    function clamp(num, min, max) {
        return num < min ? min : num > max ? max : num;
    }

	function setPos(event) {
		const { top, bottom, left, right } = refs.container.getBoundingClientRect();
		const extents = type === 'vertical' ? [top, bottom] : [left, right];
		const px = clamp(
			type === 'vertical' ? event.clientY : event.clientX,
			extents[0] + min,
			extents[1] - min
		);
		pos = 100 * (px - extents[0]) / (extents[1] - extents[0]);
		dispatch('change');
	}
	function drag(node, callback) {
		const mousedown = event => {
			if (event.which !== 1) return;
			event.preventDefault();
			dragging = true;
			const onmouseup = () => {
				dragging = false;
				window.removeEventListener('mousemove', callback, false);
				window.removeEventListener('mouseup', onmouseup, false);
			};
			window.addEventListener('mousemove', callback, false);
			window.addEventListener('mouseup', onmouseup, false);
		}
		node.addEventListener('mousedown', mousedown, false);
		return {
			destroy() {
				node.removeEventListener('mousedown', onmousedown, false);
			}
		};
	}
	$: side = type === 'horizontal' ? 'left' : 'top';
	$: dimension = type === 'horizontal' ? 'width' : 'height';
</script>

<style>
	.container {
		position: relative;
		width: 100%;
		height: 100%;
	}
	.pane {
		position: relative;
		float: left;
		width: 100%;
		height: 100%;
    overflow: hidden;
	}
	.mousecatcher {
		position: absolute;
		left: 0;
		top: 0;
		width: 100%;
		height: 100%;
		/* background: rgba(255,255,255,.01); */
	}
	.divider {
		position: absolute;
		z-index: 10;
		display: none;
	}
	.divider::after {
		content: '';
		position: absolute;
		/* background-color: #eee; */
		background-color: var(--border-color);
	}
	.horizontal {
		padding: 0 8px;
		width: 4px;
		height: 100%;
		cursor: ew-resize;
	}
	.horizontal::after {
		left: 8px;
		top: 0;
		width: 4px;
		height: 100%;
	}
	.vertical {
		padding: 8px 0;
		width: 100%;
		height: 0;
		cursor: ns-resize;
	}
	.vertical::after {
		top: 8px;
		left: 0;
		width: 100%;
		height: 4px;
	}
	.left, .right, .divider {
		display: block;
	}
	.left, .right {
		height: 100%;
		float: left;
	}
	.top, .bottom {
		position: absolute;
		width: 100%;
	}
	.top { top: 0; }
	.bottom { bottom: 0; }
</style>

<div class="container" bind:this={refs.container}>
	<div class="pane" style="{dimension}: {pos}%;">
		<slot name="a"></slot>
	</div>

	<div class="pane" style="{dimension}: {100 - (pos)}%;">
		<slot name="b"></slot>
	</div>

	{#if !fixed}
		<div class="{type} divider" style="{side}: calc({pos}% - 10px)" use:drag={setPos}></div>
	{/if}
    <slot />
</div>

{#if dragging}
	<div class="mousecatcher"></div>
{/if}
