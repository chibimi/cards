<template>
	<div id="app" class="container mt-3">
		<Selector v-on:select_card="setCardID" v-on:change_faction="setFaction" v-on:change_category="setCategory" v-on:change_language="setLanguage"/>
		<div class="row">
			<Card v-if="cardID !== null" :selected="cardID" :faction="faction" :category="category" :key="componentKey"/>
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
			componentKey: 0,
			cardID: null,
			faction: null,
			category:null,
		};
	},
	methods: {
		setCardID: function(cardID) {
			if (cardID === 0&& this.cardID===0) {
				this.cardID = -1
			} else {
				this.cardID = cardID
			}
		},
		setLanguage: function(language) {
			if (language === this.$language) {
				return
			}
			this.$change_language(language)
			this.forceRerender()
		},
		setFaction: function(faction) {
			this.faction = faction
		},
		setCategory: function(category) {
			this.category = category
		},
		forceRerender() {
			this.componentKey += 1;  
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
