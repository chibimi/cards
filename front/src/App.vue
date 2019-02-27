<template>
	<div id="app" class="container">
		<Selector v-on:selected="setCardID"/>
		<div class="row">
			<Card v-if="cardID !== null" :id="cardID" :factions="factions" :categories="categories"/>
		</div>
	</div>
</template>

<script>
import Card from "./components/card.vue"
import Selector from "./components/selector.vue"
export default {
	name: "app",
	components: {
		Card, Selector
	},
	data() {
		return {
			cardID: null
		};
	},
	methods: {
		setCardID: function(cardID){
			this.cardID=cardID;
		}
	}
};
</script>

<style>
#app {
	font-family: "Avenir", Helvetica, Arial, sans-serif;
	-webkit-font-smoothing: antialiased;
	-moz-osx-font-smoothing: grayscale;
	text-align: center;
	color: #2c3e50;
}

body,
html {
	height: 100%;
}

.h-head {
	height: 10vh;
}
.h-content {
	height: 90vh;
}
</style>
