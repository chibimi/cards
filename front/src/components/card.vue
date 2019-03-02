<template>
	<div class="w-100">
		<nav>
			<div class="nav nav-tabs" id="nav-tab" role="tablist">
				<a class="nav-item nav-link active" id="nav-ref-tab" data-toggle="tab" href="#nav-ref" role="tab" aria-controls="nav-ref" aria-selected="true">Ref</a>
				<a
					v-if="id>0"
					class="nav-item nav-link"
					id="nav-models-tab"
					data-toggle="tab"
					href="#nav-models"
					role="tab"
					aria-controls="nav-models"
					aria-selected="false"
				>Models</a>
				<a
					class="nav-item nav-link"
					id="nav-abilities-tab"
					data-toggle="tab"
					href="#nav-abilities"
					role="tab"
					aria-controls="nav-abilities"
					aria-selected="false"
				>Abilities</a>
				<a
					class="nav-item nav-link"
					id="nav-spells-tab"
					data-toggle="tab"
					href="#nav-spells"
					role="tab"
					aria-controls="nav-spells"
					aria-selected="false"
				>Spells</a>
				<a class="nav-item nav-link" id="nav-feat-tab" data-toggle="tab" href="#nav-feat" role="tab" aria-controls="nav-feat" aria-selected="false">Feat</a>
			</div>
		</nav>
		<div class="tab-content" id="nav-tabContent">
			<div class="tab-pane fade show active" id="nav-ref" role="tabpanel" aria-labelledby="nav-ref-tab">
				<Ref :id="id" :faction="faction" :category="category" v-on:new_card="newCard" v-on:remove_card="removeCard"></Ref>
			</div>
			<div class="tab-pane fade" id="nav-models" role="tabpanel" aria-labelledby="nav-models-tab">
				<Models :id="id"></Models>
			</div>
			<div class="tab-pane fade" id="nav-abilities" role="tabpanel" aria-labelledby="nav-abilities-tab">abilities</div>
			<div class="tab-pane fade" id="nav-spells" role="tabpanel" aria-labelledby="nav-spells-tab">spells</div>
			<div class="tab-pane fade" id="nav-feat" role="tabpanel" aria-labelledby="nav-feat-tab">feat</div>
		</div>
	</div>
</template>

<script>
// import Ability from "./ability.vue";
import Models from "./models.vue";
// import Spell from "./spell.vue";
// import Feat from "./feat.vue";
import Ref from "./ref.vue";
export default {
	name: "Card",
	props: ["selected", "faction", "category"],
	components: { Ref, Models },
	watch: {
		selected: function(newVal) {
			this.id = newVal
		}
	},
	created: function() {
		this.id=this.selected;
	},
	data() {
		return {
			id:0,
		};
	},
	methods: {
		newCard: function(cardID){
			this.id=cardID;
		},
		removeCard: function(){
			this.id=0;
		}
	},
};
</script>

<style scoped>
</style>
