<template>
	<div class="w-100">
		<div class="row px-3">

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
				class="col-2 mt-1"
			></v-autocomplete>
			<input v-model="selectedAbility.original_name" type="text" class="form-control col-2" placeholder="English Name">
			<input v-model="selectedAbility.name" type="text" class="form-control col-2" placeholder="French Name">
			<input v-model="selectedAbility.cost" type="text" class="form-control col-1" placeholder="spd">
			<input v-model="selectedAbility.rng" type="text" class="form-control col-1" placeholder="spd">
			<input v-model="selectedAbility.aoe" type="text" class="form-control col-1" placeholder="spd">
			<input v-model="selectedAbility.pow" type="text" class="form-control col-1" placeholder="spd">
			<input v-model="selectedAbility.dur" type="text" class="form-control col-1" placeholder="spd">
			<input v-model="selectedAbility.off" type="text" class="form-control col-1" placeholder="spd">
		</div>

		<div class="row px-3 mt-2">
			<textarea v-model="selectedAbility.description" type="text" class="form-control col-11" rows="3" placeholder/>
			<div class="col-1 px-0">
				<button v-if="spell.id || selectedAbility.id" type="submit" class="form-control btn btn-success" @click="save(selectedAbility)">Update</button>
				<button v-if="spell.id" type="submit" class="form-control btn btn-danger" @click="remove(selectedAbility)">Delete</button>
				<button v-if="!spell.id && selectedAbility.id" type="submit" class="form-control btn btn-primary" @click="add(selectedAbility)">Add</button>
				<button v-if="!spell.id && !selectedAbility.id" type="submit" class="form-control btn btn-primary" @click="save(selectedAbility)">Add</button>
			</div>
		</div>
	</div>
</template>

<script>
import ItemTemplate from "./ItemTemplate.vue";
export default {
	name: "Spell",
	props: ["spellsList", "spell"],
	watch: {
		spell: function(newVal, oldVal) {
			console.log("CHANGED ABILITY", newVal, oldVal);
			this.selectedAbility = newVal;
		}
	},
	created: function() {
		this.selectedAbility = this.spell;
	},
	data() {
		return {
			selectedAbility: {},
			template: ItemTemplate,
			items: []
		};
	},
	methods: {
		save: function(spell) {
			if (spell.id == null) {
				spell.id = 0;
			}
			this.$http
				.put("http://localhost:9901/spells/" + spell.id, spell)
				.then(function(res) {
					console.log(res);
					if (res.status === 201) {
						spell.id = res.data;
						this.add(spell);
					}
				});
		},
		add: function(spell) {
			this.$emit("add", spell);
		},
		remove: function() {
			console.log("remove");
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
			console.log("SELECTED", item, this.spell);
			this.selectedAbility = item;
			// this.spell.id = item.id;
			// this.selectedAbility.name = item.name;
			// this.selectedAbility.original_name = item.original_name;
			// this.selectedAbility.type = item.type;
			// this.selectedAbility.description = item.description;
		}
	}
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
