<template>
	<div class="w-100">
		<div v-if="update">
			<div class="form-group row">
				<div class="col-6">
					<div class="row">
						<label class="col-form-label col-3">English Name</label>
						<input v-model="card.title" type="text" class="form-control col-8" placeholder="English Name">
					</div>
					<div class="row">
						<label class="col-form-label col-3">Name</label>
						<input v-model="card.name" type="text" class="form-control col-8" placeholder="Translated Full name">
					</div>
					<div class="row">
						<label class="col-form-label col-3">Type <Tooltip :txt="help.type"/></label>
						<input v-model="card.properties" type="text" class="form-control col-8" placeholder="Translated Type">
					</div>
					<div class="row">
						<label class="col-form-label col-3">Faction</label>
						<select v-model="card.faction_id" class="form-control col-8">
							<option v-for="f in factions" :key="f.id" v-bind:value="f.id">{{f.name}}</option>
						</select>
					</div>
					<div class="row">
						<label class="col-form-label col-3">Type</label>
						<select v-model="card.category_id" class="form-control col-8">
							<option v-for="c in categories" :key="c.id" :value="c.id">{{c.name}}</option>
						</select>
					</div>
				</div>

				<div class="col-6">
					<div class="row">
						<label class="col-form-label col-3">FA</label>
						<input v-model="card.fa" type="text" class="form-control col-2">
					</div>
					<div class="row">
						<label class="col-form-label col-3">Cost</label>
						<input v-model="card.cost" type="text" class="form-control col-2">
						<label v-if="card.category_id===5" class="col-form-label col-4">Cost max</label>
						<input v-if="card.category_id===5" v-model="card.cost_max" type="text" class="form-control col-2">
					</div>
					<div class="row">
						<label class="col-form-label col-3">Nb model</label>
						<input v-model="card.models_cnt" type="text" class="form-control col-2">
						<label v-if="card.category_id===5" class="col-form-label col-4">Nb model max</label>
						<input v-if="card.category_id===5" v-model="card.models_max" type="text" class="form-control col-2" placeholder>
					</div>
					<div class="row">
						<label class="col-form-label col-3">Main ID <Tooltip :txt="help.main_id"/></label>
						<input v-model="card.main_card_id" type="text" class="form-control col-2">
					</div>
				</div>
			</div>
		</div>

		<div v-else class="text-left row">
			<div class="col-4">
				<div>
					<b>Name:</b>
					{{card.name}}
				</div>
				<div>
					<b>Tag:</b>
					{{card.properties}}
				</div>
				<div>
					<b>Faction:</b>
					{{card.faction_id}}
				</div>
				<div>
					<b>Type:</b>
					{{card.category_id}}
				</div>
			</div>
			<div class="col-4">
				<div>
					<b>FA:</b>
					{{card.fa}}
				</div>
				<div>
					<b>Cost:</b>
					{{card.cost}} for {{card.models_cnt}} models
				</div>
				<div v-if="card.category_id===5">
					<b>Cost max:</b>
					{{card.cost_max}} for {{card.models_max}} models
				</div>
				<div>
					<b>Main card ID:</b>
					{{card.main_card_id}}
				</div>
				<div>
					<b>Status:</b>
					{{card.status}}
				</div>
			</div>
		</div>
	</div>
</template>

<script>
import { Factions, Categories } from "./const.js";
import Tooltip from "./tooltip.vue";
import { EventBus } from '../main.js';

export default {
	name: "Card",
	props: ["ref_id", "ref_status"],
	components: {Tooltip},
	watch: {
		ref_id: function(newVal) {
			this.get(newVal);
		},
		ref_status: function(newVal) {
			this.card.status = newVal;
			this.update = !(this.card.id > 0 && this.card.status !== "wip")
		}
	},
	created: function() {
		this.get(this.ref_id);
	},
	mounted: function(){
		EventBus.$on('mega_save', () => {
			this.save(this.card)
		})
	},
	data() {
		return {
			factions: Factions,
			categories: Categories,
			card: {},
			vo: {},
			update: true,
			alert: "",
			alert_success: false,
			help: {
				main_id: "ID of the secondary card in case of reference having 2 distinct models like Bethayne & Belphagor. Main ID can be found in the model selector after the # (#ID)",
				type: "Tags just under the card name. Example: 'Blighted Nyss Unit'",
			}
		};
	},
	methods: {
		get: function(refID) {
			this.$http
				.get(process.env.VUE_APP_API_ENDPOINT+ "/ref/" + refID + "?lang=" + this.$language)
				.then(function(res) {
					console.log(res);
					this.card = res.data;
					if (this.card.id > 0 && this.card.status !== "wip") {
						this.update = false;
					}
				});
			this.$http
				.get(process.env.VUE_APP_API_ENDPOINT+ "/ref/" + refID + "/vo")
				.then(function(res) {
					console.log(res);
					this.vo = res.data;
				});
		},
		save: function(card) {
			if (card.id == null) {
				return
			}
			this.alert = ""
			this.$http
				.put(process.env.VUE_APP_API_ENDPOINT+ "/ref/" + card.id + "?lang=" + this.$language, card)
				.then(function(res) {
					console.log(res);
					this.alert = "save success"
					this.alert_success = true
					if (this.card.id > 0 && this.card.status !== "wip") {
						this.update = false;
					}
					EventBus.$emit('refresh_selector', card.id);
				}).catch(function(err){
					this.alert = "error: "+err.data
					this.alert_success = false
				});
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
