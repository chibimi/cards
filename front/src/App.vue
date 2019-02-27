<template>
	<div id="app" class="container">
		<Selector v-on:selected="setCardID"/>
		<div class="row">
			<Card v-if="cardID !== null" :id="cardID" :abilitiesList="abilitiesList"/>
		</div>
	</div>
</template>

<script>
import Card from "./components/card.vue";
import Selector from "./components/selector.vue";
export default {
	name: "app",
	components: {
		Card,
		Selector
	},
	data() {
		return {
			cardID: null,
			abilitiesList: []
		};
	},
	created: function() {
		this.$http.get("http://localhost:9901/abilities").then(function(res) {
			console.log(res);
			this.abilitiesList = res.data;
		});
	},
	methods: {
		setCardID: function(cardID) {
			if (cardID === 0) {
				this.cardID = -1 - this.cardID;
			} else {
				this.cardID = cardID;
			}
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
