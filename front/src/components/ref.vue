<template>
	<div class="w-100 mt-4">
		<nav>
			<div class="nav nav-tabs" id="nav-tab" role="tablist">
				<a 
					class="nav-item nav-link active" 
					id="nav-ref-tab" 
					data-toggle="tab" 
					href="#nav-ref" 
					role="tab" 
					aria-controls="nav-ref" 
					aria-selected="true"
				>Ref</a>
				<a
					v-if="reference.id>0"
					class="nav-item nav-link"
					id="nav-models-tab"
					data-toggle="tab"
					href="#nav-models"
					role="tab"
					aria-controls="nav-models"
					aria-selected="false"
				>Models</a>
				<!--
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
				-->
				<a
					v-if="reference.id>0"
					class="nav-item nav-link"
					id="nav-spells-tab"
					data-toggle="tab"
					href="#nav-spells"
					role="tab"
					aria-controls="nav-spells"
					aria-selected="false"
				>Spells & Animus</a> 
				<a
					v-if="reference.id>0 && (reference.category_id=== 1 || reference.category_id===2)"
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
				<Card :ref_id="reference.id" />
			</div>
			<div class="tab-pane fade" id="nav-models" role="tabpanel" aria-labelledby="nav-models-tab">
				<Models v-if="reference.id>0" :ref_id="reference.id" />
			</div>
			<!-- <div class="tab-pane fade" id="nav-abilities" role="tabpanel" aria-labelledby="nav-abilities-tab">
				<Abilities v-if="card.id>0" :card="card"></Abilities>
			</div>-->
			<div class="tab-pane fade" id="nav-spells" role="tabpanel" aria-labelledby="nav-spells-tab">
				<Spells v-if="reference.id>0" :id="reference.id"></Spells>
			</div> 
			<div class="tab-pane fade" id="nav-feat" role="tabpanel" aria-labelledby="nav-feat-tab">
				<Feat v-if="reference.id>0 && (reference.category_id=== 1 || reference.category_id===2)" :id="reference.id"></Feat>
			</div>
		</div>
	</div>
</template>

<script>
// import Abilities from "./abilities.vue";
import Models from "./models.vue";
import Spells from "./spells.vue";
import Feat from "./feat.vue";
import Card from "./card.vue";
export default {
	name: "Ref",
	props: ["reference"],
	components: { Card, Models, Feat, Spells},
	// components: { Card, Models, Abilities, Spells, Feat },
};
</script>

<style scoped>
</style>
