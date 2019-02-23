<template>
	<div id="app">
		<b-container>
			<b-row class="h-head">
				<b-col><Header name="Emilie"/></b-col>
			</b-row>
			<b-row class="h-content">
				<b-col><Content/></b-col>
			</b-row>
		</b-container>
	</div>
</template>

<script>

import Header from './components/header.vue'
import Content from './components/content.vue'

export default {
	name: 'app',
	components: {
		Header,
		Content,
	}
}
</script>

<style>
#app {
	font-family: 'Avenir', Helvetica, Arial, sans-serif;
	-webkit-font-smoothing: antialiased;
	-moz-osx-font-smoothing: grayscale;
	text-align: center;
	color: #2c3e50;
}

body, html {
  height:100%;
}

.h-head{
	height:5vh;
}
.h-content{
	height:95vh;
}
</style>
