<template>
	<div class="w-100">
		<div v-if="!update" class="row px-3">
			<span class="col-2 text-left">{{selectedAbility.name}} ({{selectedAbility.original_name}})</span>
			<span class="col-2 text-left">
				Cost: {{selectedAbility.cost}}, Range: {{selectedAbility.rng}}<br>
				AoE: {{selectedAbility.aoe}}, Pow: {{selectedAbility.pow}}<br>
				Dur: {{selectedAbility.dur}}, Off: {{selectedAbility.off}}</span>
			<span class="col-7 text-left">{{selectedAbility.description}}</span>
			<span class="col-1 form-inline text-right">
				<button type="submit" class="btn-sm btn-success" @click="update = true">U</button>
				<button type="submit" class="btn-sm btn-danger" @click="$emit('remove')">X</button>
			</span>
		</div>

		<div v-if="update" class="row px-3">
			<span v-if="!spell.id" class="col-2 text-left"></span>
			<span class="col-2 text-left">Original Name</span>
			<span class="col-2 text-left">Name</span>
			<span class="col-1 text-left">Cost</span>
			<span class="col-1 text-left">Range</span>
			<span class="col-1 text-left">Aoe</span>
			<span class="col-1 text-left">Pow</span>
			<span class="col-1 text-left">Dur</span>
			<span class="col-1 text-left">Off</span>
			<span v-if="spell.id" class="col-2 text-left"></span>
			<v-autocomplete
				v-if="!spell.id"
				:items="items"
				:get-label="getLabel"
				@update-items="updateItems"
				:component-item="template"
				:auto-select-one-item="false"
				@item-selected="selectedItem"
				@input="inputItem"
				class="col-2 mt-1"
			></v-autocomplete>
			<input v-if="!selectedAbility.id" v-model="selectedAbility.original_name" type="text" class="form-control col-2" placeholder="English Name">
			<label v-if="selectedAbility.id" class="col-form-label col-2 text-left">{{selectedAbility.original_name}}</label>
			<input v-model="selectedAbility.name" type="text" class="form-control col-2" placeholder="French Name">
			<input v-model="selectedAbility.cost" type="text" class="form-control col-1" placeholder="cost">
			<input v-model="selectedAbility.rng" type="text" class="form-control col-1" placeholder="rng">
			<input v-model="selectedAbility.aoe" type="text" class="form-control col-1" placeholder="aoe">
			<input v-model="selectedAbility.pow" type="text" class="form-control col-1" placeholder="pow">
			<input v-model="selectedAbility.dur" type="text" class="form-control col-1" placeholder="dur">
			<input v-model="selectedAbility.off" type="text" class="form-control col-1" placeholder="off">
			<textarea v-model="selectedAbility.description" type="text" class="form-control col-11" rows="3" placeholder/>
			<div class="col-1 px-0">
				<button v-if="spell.id || selectedAbility.id" type="submit" class="form-control btn btn-success" @click="save(selectedAbility)">Update</button>
				<button v-if="spell.id" type="submit" class="form-control btn btn-danger" @click="remove(selectedAbility)">Delete</button>
				<button v-if="!spell.id && selectedAbility.id" type="submit" class="form-control btn btn-primary" @click="add(selectedAbility)">Add</button>
				<button v-if="!spell.id && !selectedAbility.id" type="submit" class="form-control btn btn-primary" @click="save(selectedAbility)">Add</button>
			</div>
		</div>	
		<hr>
	</div>
</template>

<script>
import ItemTemplate from "./ItemTemplate.vue";
export default {
	name: "Spell",
	props: ["spellsList", "spell"],
	watch: {
		spell: function(newVal) {
			this.selectedAbility = newVal;
			if (!this.spell.id){
				this.update=true;
			}
		}
	},
	created: function() {
		this.selectedAbility = this.spell;
		if (!this.spell.id){
			this.update=true;
		}
	},
	data() {
		return {
			selectedAbility: {},
			template: ItemTemplate,
			items: [],
			update: false
		};
	},
	methods: {
		save: function(spell) {
			if (spell.id == null) {
				spell.id = 0;
			}
			this.$http
				.put(process.env.VUE_APP_API_ENDPOINT+ "/spells/" + spell.id, spell)
				.then(function(res) {
					console.log(res);
					if (this.spell.id > 0 && this.selectedAbility.id > 0) {
						this.update=false;
					}
					if (res.status === 201) {
						spell.id = res.data;
						this.add(spell);
						this.new(spell);
					} else if (res.status === 200) {
						this.updateSpell();
					}	
				});
		},
		add: function(spell) {
			this.$emit("add", spell);
		},
		new: function(spell) {
			this.$emit("new", spell);
		},
		updateSpell: function() {
			this.$emit("update");
		},
		remove: function() {
			this.$emit("remove");
		},
		getLabel: function(item) {
			if (!item) {
				return;
			}
			return item.original_name;
		},
		updateItems(text) {
			this.items = this.spellsList.filter(item =>
				item.original_name.toLowerCase().startsWith(text.toLowerCase())
			);
		},
		selectedItem(item) {
			this.selectedAbility = item;
		},
		inputItem(item) {
			if (item === null){
				this.selectedAbility = {}
			}
		}
	}
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
