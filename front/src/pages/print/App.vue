<template>
	<div id="app">
		<Selector  />
	</div>
</template>

<script>
import Selector from './components/selector.vue'

export default {
	name: 'app',
	components: { Selector },
	methods: {},
}
</script>

<style lang="scss">
@import 'custom.scss';

#app {
	@extend .container;
	font-family: 'Avenir', Helvetica, Arial, sans-serif;
	-webkit-font-smoothing: antialiased;
	-moz-osx-font-smoothing: grayscale;
	color: #2c3e50;
}
::-webkit-input-placeholder {
	color: #aeaeae !important;
	font-style: italic !important;
}
</style>
