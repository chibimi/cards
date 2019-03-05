<template>
	<div class="w-100 mt-4">
		<nav>
			<div class="nav nav-tabs" id="nav-tab" role="tablist">
				<a class="nav-item nav-link active" id="nav-ref-tab" data-toggle="tab" href="#nav-ref" role="tab" aria-controls="nav-ref" aria-selected="true">Ref</a>
				<a
					v-if="card.id>0"
					class="nav-item nav-link"
					id="nav-models-tab"
					data-toggle="tab"
					href="#nav-models"
					role="tab"
					aria-controls="nav-models"
					aria-selected="false"
				>Models</a>
				<a
					v-if="card.id>0"
					class="nav-item nav-link"
					id="nav-abilities-tab"
					data-toggle="tab"
					href="#nav-abilities"
					role="tab"
					aria-controls="nav-abilities"
					aria-selected="false"
				>Abilities</a>
				<a
					v-if="card.id>0"
					class="nav-item nav-link"
					id="nav-spells-tab"
					data-toggle="tab"
					href="#nav-spells"
					role="tab"
					aria-controls="nav-spells"
					aria-selected="false"
				>Spells & Animus</a>
				<a
					v-if="card.id>0 && (card.category_id=== 1 || card.category_id===2)"
					class="nav-item nav-link"
					id="nav-feat-tab"
					data-toggle="tab"
					href="#nav-feat"
					role="tab"
					aria-controls="nav-feat"
					aria-selected="false"
				>Feat</a>
			</div>
		</nav>
		<div class="tab-content" id="nav-tabContent">
			<div class="tab-pane fade show active" id="nav-ref" role="tabpanel" aria-labelledby="nav-ref-tab">
				<Ref :selectedCard="card" :faction="faction" :category="category" v-on:new_card="newCard" v-on:remove_card="removeCard"></Ref>
			</div>
			<div class="tab-pane fade" id="nav-models" role="tabpanel" aria-labelledby="nav-models-tab">
				<Models
					v-if="card.id>0"
					:id="card.id"
					:models="card.models"
					v-on:add="addModel"
					v-on:remove="removeModel"
					v-on:remove_weapon="removeWeapon"
					v-on:add_weapon="addWeapon"
				></Models>
			</div>
			<div class="tab-pane fade" id="nav-abilities" role="tabpanel" aria-labelledby="nav-abilities-tab">
				<Abilities v-if="card.id>0" :card="card"></Abilities>
			</div>
			<div class="tab-pane fade" id="nav-spells" role="tabpanel" aria-labelledby="nav-spells-tab">
				<Spells v-if="card.id>0" :id="card.id"></Spells>
			</div>
			<div class="tab-pane fade" id="nav-feat" role="tabpanel" aria-labelledby="nav-feat-tab">
				<Feat v-if="card.id>0 && (card.category_id=== 1 || card.category_id===2)" :id="card.id"></Feat>
			</div>
		</div>
	</div>
</template>

<script>
import Abilities from "./abilities.vue";
import Models from "./models.vue";
import Spells from "./spells.vue";
import Feat from "./feat.vue";
import Ref from "./ref.vue";
export default {
	name: "Card",
	props: ["selected", "faction", "category"],
	components: { Ref, Models, Abilities, Spells, Feat },
	watch: {
		selected: function(newVal) {
			this.getCard(newVal);
		}
	},
	created: function() {
		this.getCard(this.selected);
	},
	data() {
		return {
			card: {
				id: 0,
				models: []
			}
		};
	},
	methods: {
		getCard: async function(cardID) {
			this.reset();
			if (cardID <=0){return;}
			this.$http
				.get("http://localhost:9901/cards/" + cardID)
				.then(function(res) {
					console.log(res);
					var temp =  res.data
					temp.models = [];
					this.card = temp;
					this.getModels(cardID);
				});
		},
		newCard: function(cardID) {
			this.getCard(cardID);
		},
		removeCard: function() {
			this.reset();
		},
		reset: function(){
			this.card= {
				id: 0,
				models: []
			}
		},
		getModels: async function(cardID) {
			this.card.models = [];
			this.$http
				.get("http://localhost:9901/cards/" + cardID + "/models")
				.then(function(res) {
					console.log(res);
					var models = res.data;
					this.addWeapons(models);
				});
		},
		addWeapons: async function(array) {
			for (const [i, item] of array.entries()) {
				await this.getWeapons(i, item);
			}
		},
		getWeapons: async function(i, model) {
			await this.$http
				.get("http://localhost:9901/models/" + model.id + "/weapons")
				.then(function(res) {
					console.log(res);
					model.weapons = res.data;
					this.card.models.splice(i, 1, model);
				});
		},
		removeWeapon: function(index, weaponIndex) {
			this.card.models[index].weapons.splice(weaponIndex, 1);
		},
		addWeapon: function(index, weapon) {
			this.card.models[index].weapons.push(weapon);
		},
		removeModel: function(index) {
			this.card.models.splice(index, 1);
		},
		addModel: function(model) {
			this.card.models.push(model);
		}
	}
};
</script>

<style scoped>
</style>
