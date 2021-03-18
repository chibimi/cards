<template>
	<div id="app">
		<p> TOTO </p>
	</div>
</template>

<script>

export default {
	name: 'app',
	data() {
		return {
			ref_id: null,
		}
	},
	methods: {},
}
</script>

<style lang="scss">

#app {
	font-family: 'Avenir', Helvetica, Arial, sans-serif;
	-webkit-font-smoothing: antialiased;
	-moz-osx-font-smoothing: grayscale;
	color: #2c3e50;
}
::-webkit-input-placeholder {
	color: #aeaeae !important;
	font-style: italic !important;
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
