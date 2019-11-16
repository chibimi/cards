<template>
	<div class="w-100 mt-4">
		<div class="row my-4">
			<h2 class="text-left col-7">{{reference.title}} <span class="font-italic h5">{{reference.properties}}</span> </h2>
			<select v-model="status" class="form-control col-2 mr-3">
				<option value="wip">WIP</option>
				<option value="tbv">A valider</option>
				<option value="done">Terminée</option>
			</select>
			<!-- <div class="col-4" >
				<div v-if="alert" class="alert py-2" v-bind:class="{ 'alert-success': alert_success,'alert-danger': !alert_success}">{{alert}}</div>
			</div> -->
			<div class="col-1 pl-0">
				<button type="submit" class="form-control btn btn-primary" @click="save()">Save</button>
			</div>
			<div class="col-1 pl-0">
				<button type="submit" class="form-control btn btn-danger" @click="remove(reference.id)">Delete</button>
			</div>
		</div>
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
				<a
					v-if="reference.id>0"
					class="nav-item nav-link"
					id="nav-abilities-tab"
					data-toggle="tab"
					href="#nav-abilities"
					role="tab"
					aria-controls="nav-abilities"
					aria-selected="false"
					v-on:click="abilitiesKey++"
				>Abilities</a>	
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
					v-if="reference.id>0 && (reference.category_id=== 1 || reference.category_id===2|| reference.category_id===10)"
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
				<Card :ref_id="reference.id" :ref_status="status" />
			</div>
			<div class="tab-pane fade" id="nav-models" role="tabpanel" aria-labelledby="nav-models-tab">
				<Models v-if="reference.id>0" :ref_id="reference.id" />
			</div>
			<div class="tab-pane fade" id="nav-abilities" role="tabpanel" aria-labelledby="nav-abilities-tab">	
				<Abilities v-if="reference.id>0" :ref_id="reference.id" :key="abilitiesKey"></Abilities>
			</div>
			<div class="tab-pane fade" id="nav-spells" role="tabpanel" aria-labelledby="nav-spells-tab">
				<Spells v-if="reference.id>0" :ref_id="reference.id"></Spells>
			</div> 
			<div class="tab-pane fade" id="nav-feat" role="tabpanel" aria-labelledby="nav-feat-tab">
				<Feat v-if="reference.id>0 && (reference.category_id=== 1 || reference.category_id===2|| reference.category_id===10)" :ref_id="reference.id"></Feat>
			</div>
		</div>
	</div>
</template>

<script>
import Abilities from "./abilities.vue";
import Models from "./models.vue";
import Spells from "./spells.vue";
import Feat from "./feat.vue";
import Card from "./card.vue";
import { EventBus } from '../main.js';
export default {
	name: "Ref",
	props: ["reference"],
	components: { Card, Models, Feat, Spells, Abilities},
	watch: {
		reference: function(newVal) {
			this.status = newVal.status;
		}
	},
	ceeated: function(){
		this.status = this.reference.status;
	},
	mounted: function(){
		this.status = this.reference.status;
		EventBus.$on('mega_save', function(){
			console.log(`SAVE REF`)
		});
	},
	data() {
		return {
			status: "wip",
			alert: "",
			alert_success: false,
			abilitiesKey: 0,
		};
	},
	methods: {
		reset: function() {
			this.alert = ""
			this.alert_success = false
		},
		save: function() {
			EventBus.$emit('mega_save');
		},
		// remove: function(card) {
		// 				this.alert = ""
		// 	this.$http
		// 		.delete(process.env.VUE_APP_API_ENDPOINT+ "/cards/" + card.id)
		// 		.then(function(res) {
		// 			console.log(res);
		// 			if (res.status === 204) {
		// 				this.$emit('remove_card', card.id)
		// 			}
		// 		}).catch(function(err){
		// 			this.alert = "error: "+err
		// 			this.alert_success = false
		// 		});
		// },
	}
};
</script>

<style scoped>
</style>
