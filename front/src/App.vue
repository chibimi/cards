<template>
	<div id="app">
		<Selector v-on:select_ref="ref_id = $event" :key="$language" />
		<Ref v-if="ref_id !== null" :ref_id="ref_id" :key="$language" />
	</div>
</template>

<script>
import Selector from './components/selector.vue'
import Ref from './components/ref.vue'

export default {
	name: 'app',
	components: { Selector, Ref },
	data() {
		return {
			ref_id: null,
		}
	},
	methods: {},
}
</script>

<style lang="scss">
@import 'custom.scss';

#app {
	@extend .container;

}
// FIXME
// #app {
// 	font-family: "Avenir", Helvetica, Arial, sans-serif;
// 	-webkit-font-smoothing: antialiased;
// 	-moz-osx-font-smoothing: grayscale;
// 	text-align: center;
// 	color: #2c3e50;
// }

// body,
// html {
// 	height: 100%;
// }
// .h-head {
// 	height: 10vh;
// }
// .h-content {
// 	height: 90vh;
// }
// ::-webkit-input-placeholder {
// 	color: #aeaeae !important;
// 	font-style: italic !important;
// }
// .vo {
// 	color: #828da9;
// }
</style>
